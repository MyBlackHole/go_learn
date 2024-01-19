package emulator

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type ContextTraceType string

const ContextTraceKey = ContextTraceType("ctx-trace-info")

type TraceCtxt struct {
	RequestRecorder  *RequestRecorder
	ResponseRecorder *ResponseRecorder

	FuncName string
	AmzReqID string
}

func getOpName(name string) (op string) {
	op = strings.TrimPrefix(name, "github.com/minio/minio/cmd.")
	op = strings.TrimSuffix(op, "Handler-fm")
	op = strings.Replace(op, "objectAPIHandlers", "s3", 1)
	op = strings.Replace(op, "adminAPIHandlers", "admin", 1)
	op = strings.Replace(op, "(*storageRESTServer)", "storageR", 1)
	op = strings.Replace(op, "(*peerRESTServer)", "peer", 1)
	op = strings.Replace(op, "(*lockRESTServer)", "lockR", 1)
	op = strings.Replace(op, "(*stsAPIHandlers)", "sts", 1)
	op = strings.Replace(op, "(*peerS3Server)", "s3", 1)
	op = strings.Replace(op, "ClusterCheckHandler", "health.Cluster", 1)
	op = strings.Replace(op, "ClusterReadCheckHandler", "health.ClusterRead", 1)
	op = strings.Replace(op, "LivenessCheckHandler", "health.Liveness", 1)
	op = strings.Replace(op, "ReadinessCheckHandler", "health.Readiness", 1)
	op = strings.Replace(op, "-fm", "", 1)
	return op
}

// http 跟踪
func httpTrace(f http.HandlerFunc, logBody bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tc, ok := r.Context().Value(ContextTraceKey).(*TraceCtxt)
		if !ok {
			// Tracing is not enabled for this request
			f.ServeHTTP(w, r)
			return
		}

		tc.FuncName = getOpName(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		tc.RequestRecorder.LogBody = logBody
		tc.ResponseRecorder.LogAllBody = logBody
		tc.ResponseRecorder.LogErrBody = true

		f.ServeHTTP(w, r)
	}
}

func httpTraceAll(f http.HandlerFunc) http.HandlerFunc {
	return httpTrace(f, true)
}

func httpTraceHdrs(f http.HandlerFunc) http.HandlerFunc {
	return httpTrace(f, false)
}
