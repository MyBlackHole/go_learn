package main

// xg响应实体
type XGResp struct {
	Code int    `json:"code"`
	Data XGData `json:"data"`
	Mes  string `json:"mes"`
}

type InfoSrc struct {
}

type Statuses struct {
	Source Blog `json:"_source"`
}
type XGData struct {
	TotalNumber int        `json:"total_number"`
	Took        int        `json:"took"`
	Statuses    []Statuses `json:"statuses"`
	Page        int        `json:"page"`
}

// 响应实体
type Resp struct {
	Code    string `json:"code"`
	Data    Data   `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Gather struct {
	SiteName   string `json:"site_name"`
	SiteDomain string `json:"site_domain"`
	Stime      int    `json:"stime"`
}
type Hashcode struct {
	Num1 int64 `json:"1"`
	Num2 int64 `json:"2"`
	Num3 int64 `json:"3"`
	Num4 int64 `json:"4"`
	Num5 int64 `json:"5"`
}
type Analysis struct {
	Sentiment int      `json:"sentiment"`
	InfoSrc   InfoSrc  `json:"info_src"`
	Hashcode  Hashcode `json:"hashcode"`
	Mentions  []string `json:"mentions"`
	Hashtag   []string `json:"hashtag"`
}
type User struct {
	UID              interface{} `json:"uid"`
	URL              string      `json:"url"`
	Name             string      `json:"name"`
	NickName         int         `json:"nickname"`
	ProfileImgURL    string      `json:"profile_img_url"`
	Gender           string      `json:"gender"`
	LangCode         string      `json:"lang_code"`
	Level            int         `json:"level"`
	Verified         int         `json:"verified"`
	CreatedAt        int         `json:"created_at"`
	VerifiedType     int         `json:"verified_type"`
	VerifiedReason   interface{} `json:"verified_reason"`
	DescriPtion      string      `json:"description"`
	Location         []string    `json:"location"`
	LocationCode     string      `json:"location_code"`
	FollowersCount   int         `json:"followers_count"`
	BiFollowersCount int         `json:"bi_followers_count"`
	StatusesCount    int         `json:"statuses_count"`
	FriendsCount     int         `json:"friends_count"`
}
type Publisher struct {
	SiteName string `json:"site_name"`
	Name     string `json:"name"`
	ID       string `json:"id"`
	Platform string `json:"platform"`
	Entity   string `json:"entity"`
}
type Blog struct {
	UUID         string      `json:"uuid"`
	Wtype        int         `json:"wtype"`
	Mid          string      `json:"mid"`
	URL          string      `json:"url"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	Ctime        float64     `json:"ctime"`
	Utime        int64       `json:"utime"`
	Channel      string      `json:"channel"`
	Device       string      `json:"device"`
	SurfaceImg   string      `json:"surface_img"`
	Pid          interface{} `json:"pid"`
	RootID       string      `json:"rootid"`
	VisitCount   int         `json:"visit_count"`
	ReplyCount   int         `json:"reply_count"`
	RepostCount  int         `json:"repost_count"`
	LikeCount    int         `json:"like_count"`
	ShareCount   int         `json:"share_count"`
	RepostSource string      `json:"repost_source"`
	MusicID      string      `json:"music_id"`
	PicUrls      []string    `json:"pic_urls"`
	VideoUrls    []string    `json:"video_urls"`
	Place        string      `json:"place"`
	OCR          string      `json:"ocr"`
	Analysis     Analysis    `json:"analysis"`
	Gather       Gather      `json:"gather"`
	User         User        `json:"user"`
	Publisher    Publisher   `json:"publisher"`
	Retweeted    *Blog       `json:"retweeted"`
}
type Data struct {
	Total int    `json:"total"`
	List  []Blog `json:"list"`
}

// 博文实体
type URunBlog struct {
	Groupname        string `json:"GroupName"`
	Account          string `json:"Account"`
	Addon            int64  `json:"AddOn"`
	Dsensitive       int    `json:"DSensitive"`
	Quoteblogid      string `json:"QuoteBlogID"`
	Quoteurl         string `json:"QuoteUrl"`
	Overseas         int    `json:"Overseas"`
	Quoteauthor      string `json:"QuoteAuthor"`
	Name             string `json:"Name"`
	Groupid          int    `json:"GroupID"`
	Transmits        int    `json:"Transmits"`
	Hashcode         string `json:"HashCode"`
	Verifiedtype     int    `json:"VerifiedType"`
	Quotepraises     int    `json:"QuotePraises"`
	Headline         int    `json:"Headline"`
	Blogid           string `json:"BlogID"`
	Route            string `json:"_route_"`
	ID               string `json:"ID"`
	Tags             string `json:"Tags"`
	Quoteuid         string `json:"QuoteUID"`
	Quoteimageurl    string `json:"QuoteImageUrl"`
	Quotetransmits   int    `json:"QuoteTransmits"`
	Quotecontent     string `json:"QuoteContent"`
	Channel          int    `json:"Channel"`
	Weibotype        string `json:"WeiboType"`
	Dsentiment       int    `json:"DSentiment"`
	Imageurl         string `json:"ImageUrl"`
	City             string `json:"City"`
	Province         string `json:"Province"`
	UID              string `json:"UID"`
	Quoteportraiturl string `json:"QuotePortraitUrl"`
	Taskname         string `json:"TaskName"`
	Content          string `json:"Content"`
	Sitetags         string `json:"SiteTags"`
	Country          int    `json:"Country"`
	Dareawords       string `json:"DAreaWords"`
	Domain           string `json:"Domain"`
	Fans             int    `json:"Fans"`
	Quoteimages      int    `json:"QuoteImages"`
	Keywords         string `json:"Keywords"`
	Quotetime        int64  `json:"QuoteTime"`
	Sex              int    `json:"Sex"`
	Time             int64  `json:"Time"`
	URL              string `json:"Url"`
	At               string `json:"At"`
	Language         int    `json:"Language"`
	Praises          int    `json:"Praises"`
	Persons          string `json:"Persons"`
	Sitetype         int    `json:"SiteType"`
	Taskid           int    `json:"TaskID"`
	Comments         int    `json:"Comments"`
	Title            string `json:"Title"`
	From             string `json:"From"`
	Layer            int    `json:"Layer"`
	Videourl         string `json:"VideoUrl"`
	Portraiturl      string `json:"PortraitUrl"`
	Dsensitivewords  string `json:"DSensitiveWords"`
	Dgarbage         int    `json:"DGarbage"`
	Quotecomments    int    `json:"QuoteComments"`
	Location         string `json:"Location"`
	Addtime          int64  `json:"AddTime"`
}

type MyData struct {
	Total int        `json:"total"`
	List  []URunBlog `json:"list"`
}

// My Resp 实体
type MyResp struct {
	ERROR string `json:"error"`
	MyData
}

// 请求参数
type Token struct {
	Token string `json:"token"`
}
type Json struct {
	Keyword    string   `json:"keyword"`
	StartTime  string   `json:"startTime"`
	EndTime    string   `json:"endTime"`
	PageIndex  int      `json:"pageIndex"`
	PageSize   int      `json:"pageSize"`
	Uids       []string `json:"uids"`
	SiteTypes  []string `json:"siteTypes"`
	Status     bool     `json:"status"`
	Original   bool     `json:"original"`
	SendStatus bool     `json:"sendStatus"`
}

// app\weibo\weixin
type Person struct {
	Json
	Token
}

// 星光
type XGPerson struct {
	Source      string    `json:"source"`
	Time        string    `json:"time"`
	Keyword     string    `json:"keyword"`
	SortOf      string    `json:"sort_of"`
	Size        int       `json:"size"`
	Page        int       `json:"page"`
	MatchFields []string  `json:"match_fields"`
	Filters     []Filters `json:"filters"`
}

type Filters struct {
	Operator string   `json:"operator"`
	Field    string   `json:"field"`
	Values   []string `json:"values"`
}

type RedisData struct {
	Time string `json:"time"`
	Json
}

type RedisResp struct {
	ERROR string      `json:"error"`
	Data  []RedisData `json:"data"`
}

// 底层收包结构
type Headers struct {
	Topic     string `json:"topic"`
	Key       string `json:"key"`
	Timestamp int64  `json:"timestamp"`
}

type Package struct {
	Headers Headers `json:"headers"`
	Body    string  `json:"body"`
}
