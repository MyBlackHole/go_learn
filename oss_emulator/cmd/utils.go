package emulator

import (
	"context"
    "time"
	"github.com/minio/mux"
	"github.com/valyala/bytebufferpool"
	"net/http"
	"net/url"
	"path"
    "strings"
)

const (
	slashSeparator     = "/"
	SlashSeparatorChar = '/'
)

type KeyVal struct {
	Key string
	Val interface{}
}

type ReqInfo struct {
	// RemoteHost   string // Client Host/IP
	Host      string // Node Host/IP
	UserAgent string // User Agent
	// DeploymentID string // x-minio-deployment-id
	RequestID  string // x-amz-request-id
	API        string // API name - GetObject PutObject NewMultipartUpload etc.
	BucketName string `json:",omitempty"` // Bucket name
	ObjectName string `json:",omitempty"` // Object name
	// VersionID    string `json:",omitempty"` // corresponding versionID for the object
	// Objects      []ObjectVersion  `json:",omitempty"` // Only set during MultiObject delete handler.
	Region   string   `json:"-"`
	Owner    bool     `json:"-"`
	AuthType string   `json:"-"`
	tags     []KeyVal // Any additional info not accommodated by above fields
	// sync.RWMutex
}

type contextKeyType string

const contextLogKey = contextKeyType("miniolog")

func SetReqInfo(ctx context.Context, req *ReqInfo) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, contextLogKey, req)
}

func GetReqInfo(ctx context.Context) *ReqInfo {
	if ctx != nil {
		r, ok := ctx.Value(contextLogKey).(*ReqInfo)
		if ok {
			return r
		}
		r = &ReqInfo{}
		return r
	}
	return nil
}

func newContext(r *http.Request, w http.ResponseWriter, api string) context.Context {
	reqID := w.Header().Get(AmzRequestID)

	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := likelyUnescapeGeneric(vars["object"], url.PathUnescape)
	reqInfo := &ReqInfo{
		// DeploymentID: globalDeploymentID(),
		RequestID: reqID,
		// RemoteHost:   handlers.GetSourceIP(r),
		Host:       getHostName(r),
		UserAgent:  r.UserAgent(),
		API:        api,
		BucketName: bucket,
		ObjectName: object,
		// VersionID:    strings.TrimSpace(r.Form.Get(xhttp.VersionID)),
	}

	return SetReqInfo(r.Context(), reqInfo)
}

func likelyUnescapeGeneric(p string, escapeFn func(string) (string, error)) string {
	ep, err := unescapeGeneric(p, escapeFn)
	if err != nil {
		return p
	}
	return ep
}

func unescapeGeneric(p string, escapeFn func(string) (string, error)) (string, error) {
	ep, err := escapeFn(p)
	if err != nil {
		return "", err
	}
	return trimLeadingSlash(ep), nil
}

func trimLeadingSlash(ep string) string {
	if len(ep) > 0 && ep[0] == '/' {
		// Path ends with '/' preserve it
		if ep[len(ep)-1] == '/' && len(ep) > 1 {
			ep = path.Clean(ep)
			ep += slashSeparator
		} else {
			ep = path.Clean(ep)
		}
		ep = ep[1:]
	}
	return ep
}

func pathJoin(elem ...string) string {
	sb := bytebufferpool.Get()
	defer func() {
		sb.Reset()
		bytebufferpool.Put(sb)
	}()

	return pathJoinBuf(sb, elem...)
}

func pathJoinBuf(dst *bytebufferpool.ByteBuffer, elem ...string) string {
	trailingSlash := len(elem) > 0 && hasSuffixByte(elem[len(elem)-1], SlashSeparatorChar)
	dst.Reset()
	added := 0
	for _, e := range elem {
		if added > 0 || e != "" {
			if added > 0 {
				dst.WriteByte(SlashSeparatorChar)
			}
			dst.WriteString(e)
			added += len(e)
		}
	}

	if pathNeedsClean(dst.Bytes()) {
		s := path.Clean(dst.String())
		if trailingSlash {
			return s + SlashSeparator
		}
		return s
	}
	if trailingSlash {
		dst.WriteByte(SlashSeparatorChar)
	}
	return dst.String()
}

func pathNeedsClean(path []byte) bool {
	if len(path) == 0 {
		return true
	}

	rooted := path[0] == '/'
	n := len(path)

	r, w := 0, 0
	if rooted {
		r, w = 1, 1
	}

	for r < n {
		switch {
		case path[r] > 127:
			// Non ascii.
			return true
		case path[r] == '/':
			// multiple / elements
			return true
		case path[r] == '.' && (r+1 == n || path[r+1] == '/'):
			// . element - assume it has to be cleaned.
			return true
		case path[r] == '.' && path[r+1] == '.' && (r+2 == n || path[r+2] == '/'):
			// .. element: remove to last / - assume it has to be cleaned.
			return true
		default:
			// real path element.
			// add slash if needed
			if rooted && w != 1 || !rooted && w != 0 {
				w++
			}
			// copy element
			for ; r < n && path[r] != '/'; r++ {
				w++
			}
			// allow one slash, not at end
			if r < n-1 && path[r] == '/' {
				r++
			}
		}
	}

	// Turn empty string into "."
	if w == 0 {
		return true
	}

	return false
}

func hasSuffixByte(s string, suffix byte) bool {
	return len(s) > 0 && s[len(s)-1] == suffix
}

func unescapePath(p string) (string, error) {
	return unescapeGeneric(p, url.PathUnescape)
}

// 判断字符串是不是以 suffix 结尾
func HasSuffix(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func UTCNow() time.Time {
	return time.Now().UTC()
}

func baseDirFromPrefix(prefix string) string {
	b := path.Dir(prefix)
	if b == "." || b == "./" || b == "/" {
		b = ""
	}
	if !strings.Contains(prefix, slashSeparator) {
		b = ""
	}
	if len(b) > 0 && !strings.HasSuffix(b, slashSeparator) {
		b += slashSeparator
	}
	return b
}

func s3EncodeName(name, encodingType string) string {
	if strings.ToLower(encodingType) == "url" {
		return s3URLEncode(name)
	}
	return name
}

func s3URLEncode(s string) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			if c == ' ' {
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	var buf [64]byte
	var t []byte

	required := len(s) + 2*hexCount
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}

	if hexCount == 0 {
		copy(t, s)
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				t[i] = '+'
			}
		}
		return string(t)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == ' ':
			t[j] = '+'
			j++
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

func shouldEscape(c byte) bool {
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	switch c {
	case '-', '_', '.', '/', '*':
		return false
	}
	return true
}
