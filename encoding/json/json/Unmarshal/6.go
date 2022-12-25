package main

import (
	"encoding/json"
	"fmt"
)

// 响应实体
type Resp struct {
	Code    string `json:"code"`
	Data    Data   `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type Gather struct {
	SiteName   string   `json:"site_name"`
	SiteDomain string   `json:"site_domain"`
	Stime      int      `json:"stime"`
	InfoFlag   []string `json:"info_flag"`
}
type InfoSrc struct {
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
	UID              interface{} 
	URL              int                                                                          `json:"url"`
	Name             string                                             `json:"name"`
	NickName         int                                                `json:"nickname"`
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
	UUID         string    `json:"uuid"`
	Wtype        int       `json:"wtype"`
	Mid          string    `json:"mid"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Ctime        int       `json:"ctime"`
	Utime        int       `json:"utime"`
	Channel      string    `json:"channel"`
	Device       string    `json:"device"`
	SurfaceImg   string    `json:"surface_img"`
	Pid          string    `json:"pid"`
	RootID       string    `json:"rootid"`
	VisitCount   int       `json:"visit_count"`
	ReplyCount   int       `json:"reply_count"`
	RepostCount  int       `json:"repost_count"`
	LikeCount    int       `json:"like_count"`
	ShareCount   int       `json:"share_count"`
	RepostSource string    `json:"repost_source"`
	MusicID      string    `json:"music_id"`
	PicUrls      []string  `json:"pic_urls"`
	VideoUrls    []string  `json:"video_urls"`
	Place        string    `json:"place"`
	OCR          string    `json:"ocr"`
	Analysis     Analysis  `json:"analysis"`
	Gather       Gather    `json:"gather"`
	User         User      `json:"user"`
	Publisher    Publisher `json:"publisher"`
	Retweeted    *Blog     `json:"retweeted"`
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
	Quotetime        int    `json:"QuoteTime"`
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
	Version          int64  `json:"_version_"`
}

func main() {
	// 	data := []byte(`
	// {
	//     "code": "10005",
	//     "data": null,
	//     "success": true,
	// 	"ok": true,
	//     "message": "exception..."
	// }
	// 		`)
	data := []byte(`{
    "code": "0",
    "data": {
        "total": 32970935,
        "list": [
            {
                "visit_count": 12,
                "wtype": 1,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "代提白条咨询",
                    "id": "weibo.com|5258132230",
                    "platform": "自媒体",
                    "entity": "代提白条咨询"
                },
                "pic_urls": [
                    "https://wx4.sinaimg.cn/large/005JQAomgy1gp1p9n1g09j30hs0b6aah.jpg"
                ],
                "gather": {
                    "site_name": "新浪长微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069684,
                    "info_flag": [
                        "04",
                        "0408"
                    ]
                },
                "title": "新疆官员喊话H＆amp;M：中国人民是不好惹的",
                "analysis": {
                    "sentiment": -1,
                    "info_src": {},
                    "hashcode": {
                        "1": 2578123342185533120,
                        "2": 2065192252050464873,
                        "3": 2375189503267018258,
                        "4": 7535712152046263581,
                        "5": 273170405911411488
                    }
                },
                "uuid": "d485deac90fb11eba40b868712de27fd",
                "user": {
                    "uid": "5258132230",
                    "name": "代提白条咨询"
                },
                "content": "​​今天（29日）上午，外交部与新疆维吾尔自治区共同举行涉疆问题新闻发布会。就“美国等国家指责新疆”“H&M等品牌抵制新疆棉花”等问题，新疆维吾尔自治区人民政府新闻发言人徐贵相一一回应。\n徐贵相：美英加欧盟“失法”“失理”“失忆”“失智”\n上周，欧盟、英国、美国、加拿大分别以所谓“新疆人权”的问题为借口，对中国相关个人和实体进行了制裁，新疆维吾尔自治区人民政府新闻发言人徐贵相表示，欧盟、美国、英国、加拿大这种做法严重违反国际法和国际关系基本准则，严重干涉中国内政，严重伤害新疆各族人民的感情。徐贵相用“失法”“失理”“失忆”“失智”来评价欧盟、美国、英国、加拿大的行为。\n集体“失法”：恶意污蔑中国的有关治疆措施，堪称人类历史上最大诬陷案。\n集体“失理”：无视新疆巨大发展，睁眼说瞎话。\n集体“失忆”：对自身罪行闭口不谈。\n集体“失智”：为插手新疆事务、干涉中国内政，他们热衷于政治操弄，沉迷于“制裁”把戏，已经到了歇斯底里的地步。\n\n美国等国家为何依然指责新疆？徐贵相：暴恐频发是他们心中“最美风景线”\n长久以来，新疆受贫困和暴恐困扰，但在中国政府的努力下，摆脱了贫困、遏制了暴恐，为什么美国等国家依然在批评和指责中国新疆？\n对此，徐贵相表示，这是因为过去暴恐频发的状态，是他们心中“最美的风景线”。现在，新疆实现了稳定繁荣，人民生活安定祥和，美西方反华势力达不到目的很焦躁，于是就想法设法指责新疆，睁着眼睛说瞎话。\n徐贵相强调：“对于这种阴暗心理，我们看得很清楚，新疆所做的一切没有什么可指责的，实践证明也是成功的，我们不吃他们这一套。我们会坚定自信地按照党中央制定的方略，坦坦荡荡、坚定不移为各族群众办好事，让人民群众过上更好的生活。这是对美西方反华势力污蔑造谣的最好回答。”\n徐贵相：西方制裁的大棒挥向新疆企业，也会砸向自己\n有媒体就“近日H&M等品牌抵制新疆棉花”一事提问，徐贵相回应称：“企业不应把经济行为政治化。”\n徐贵相表示，中国早已不是1840年的中国，中国人民遭受西方列强霸权、霸凌的时代一去不复返了，包括新疆各族人民在内的全体中国人民是不好惹的。制裁的大棒挥向新疆企业的时候，也会砸向自己。希望更多类似H&M的企业，要擦亮眼睛，明辨是非。\n​\n本文部分内容转载于网络，如有侵权，麻烦联系删除！​​​​",
                "url": "http://weibo.com/ttarticle/p/show?id=2309404620387337240843"
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387143650717",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069605,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": -1,
                    "info_src": {},
                    "hashcode": {
                        "1": 2461748307155412889,
                        "2": 1380775079577525954,
                        "3": 4300048763081053186,
                        "4": 3153396924974952671,
                        "5": 1606723427413810891
                    },
                    "mentions": [
                        "微丢"
                    ],
                    "hashtag": [
                        "赵立坚问美方我们是强迫机器劳动了吗"
                    ]
                },
                "reply_count": 0,
                "uuid": "a482877890fb11ebb924d200e99237a2",
                "content": "【赵立坚：编故事者难道是说新疆强迫机器劳动？】#赵立坚问美方我们是强迫机器劳动了吗# 新疆的采棉机器怕是要说：哈？我被强迫劳动了？在29日的外交部例行记者会上，有媒体问到涉华谣言大户美国国务院针对中国网民自发抵制H&M等企业一事，美国国务院再次成为了涉疆谎言的背后推手和幕后主使，并大言不惭地称中国政府利用外国企业对华依赖打压“言论自由”。对此，赵立坚表示，2020年新疆棉花机械采摘率已达70%，这个时候一小撮国家和个人还在编造所谓“强迫劳动”的故事。既然该协会（BCI）没有发现“强迫劳动”，那么“强迫劳动”的结论是怎么来的？请你们拿出确凿证据来！ @微丢 微丢的微博视频",
                "url": "http://weibo.com/1893801487/K8ymufjIx",
                "wtype": 1,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "新浪微博",
                    "name": "微天下",
                    "id": "weibo.com|1893801487",
                    "platform": "自媒体",
                    "entity": "微天下"
                },
                "video_urls": [
                    "https://video.weibo.com/show?fid=1034:4620130005745739"
                ],
                "device": "微博发布平台专业版",
                "user": {
                    "friends_count": 1181,
                    "profile_img_url": "http://tvax4.sinaimg.cn/crop.1.0.497.497.50/70e11e0fly8gdigchaorcj20dv0dvmyc.jpg",
                    "gender": "m",
                    "level": 48,
                    "verified": 2,
                    "description": "24小时全球热点播报，新浪新闻出品。",
                    "created_at": 1292403944,
                    "verified_type": 3,
                    "uid": 1893801487,
                    "lang_code": "zh-cn",
                    "location_code": "116357",
                    "bi_followers_count": 534,
                    "statuses_count": 168466,
                    "followers_count": 30043397,
                    "name": "微天下",
                    "verified_reason": "24小时全球热点",
                    "location": [
                        "北京",
                        "东城区"
                    ]
                }
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387143388239",
                "pid": "4618389982546373",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069603,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 506260494671599246,
                        "2": 2523510537117804249,
                        "3": 4445489859500169422,
                        "4": 7407775761963770821,
                        "5": -7407775761963770821
                    },
                    "mentions": [
                        "YKYBtao-黄子韬"
                    ],
                    "hashtag": [
                        "我支持新疆棉花"
                    ]
                },
                "reply_count": 0,
                "uuid": "a3d82cc490fb11eba2589e05269dec8f",
                "content": "支持！//@YKYBtao-黄子韬:国家的利益高于一切！支持我们的新疆棉花",
                "url": "http://weibo.com/2040947827/K8ymuedr1",
                "retweeted": {
                    "repost_count": 37465583,
                    "wtype": 1,
                    "like_count": 0,
                    "mid": "4618389982546373",
                    "ctime": 1616593440,
                    "pic_urls": [
                        "http://wx2.sinaimg.cn/thumbnail/0033ImPzly1govbvfxn40j60u01hcn7u02.jpg"
                    ],
                    "reply_count": 145341,
                    "user": {
                        "friends_count": 3056,
                        "uid": 2803301701,
                        "gender": "m",
                        "level": 48,
                        "bi_followers_count": 0,
                        "statuses_count": 130418,
                        "followers_count": 129085181,
                        "verified": 2,
                        "name": "人民日报",
                        "location": [
                            "北京"
                        ],
                        "verified_type": 3
                    },
                    "content": "[话筒]#我支持新疆棉花# ​​",
                    "url": "http://weibo.com/2803301701/K7IpgaGqx"
                },
                "wtype": 2,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "闲来无事生非233",
                    "id": "weibo.com|2040947827",
                    "platform": "自媒体",
                    "entity": "闲来无事生非233"
                },
                "device": "荣耀20",
                "user": {
                    "friends_count": 736,
                    "profile_img_url": "http://tvax3.sinaimg.cn/crop.0.0.986.986.50/79a66473ly8go3orsdvlxj20re0redji.jpg",
                    "gender": "f",
                    "level": 9,
                    "verified": 0,
                    "created_at": 1300866035,
                    "verified_type": -1,
                    "uid": 2040947827,
                    "lang_code": "zh-cn",
                    "location_code": "69196",
                    "bi_followers_count": 13,
                    "statuses_count": 9470,
                    "followers_count": 193,
                    "name": "闲来无事生非233",
                    "location": [
                        "四川省",
                        "成都市"
                    ]
                }
            },
            {
                "visit_count": 244,
                "wtype": 1,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "今日影评Mtalk",
                    "id": "weibo.com|5828293600",
                    "platform": "自媒体",
                    "entity": "今日影评Mtalk"
                },
                "gather": {
                    "site_name": "新浪长微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069635,
                    "info_flag": [
                        "04",
                        "0408"
                    ]
                },
                "title": "明星代言的正确打开方式",
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 2814346738786195260,
                        "2": 831680002944435116,
                        "3": 4370595411559398674,
                        "4": 5700693515003824770,
                        "5": 6804291605133264039
                    },
                    "hashtag": [
                        "我支持新疆棉花"
                    ]
                },
                "uuid": "b760abfe90fb11ebbbe4ee68f8197055",
                "user": {
                    "uid": "5828293600",
                    "name": "今日影评Mtalk"
                },
                "content": "在前几天的活动中，曾黎在出席活动时特意佩戴了一朵新疆棉花，在采访时，主持人询问曾黎对中戏200年美女称号怎么看，曾黎则是高情商地表示：“不如我手上的新疆棉花好看。”在微博上也是发出了自己佩戴新疆棉花的照片，并配文“我爱我的中国棉 #我支持新疆棉花#”。这样情商又高三观又正的美女谁会不爱呢，给曾黎点赞👍。除了曾黎以外，还有很多明星艺人在通过自己的方式表达对新疆棉花的支持：薇娅直接下架了涉事企业的产品，而且去新疆给新疆棉花直播带货；陈建斌老师写诗支持新疆棉花，然而在因为新疆棉花引发了一系列关注度与影响力颇高的事件，最让人瞩目的，还是众多品牌代言人齐刷刷的和这些不尊重我们国家的品牌发表了解约声明。\n\n“中国艺人的态度其实是代表着我们自己的态度，在面对这种大是大非的选择上，我们每一个中国人其实是都是和国家在一起的，而且在这件事情上我们可以看到艺术家就我们的艺人他其实也是推动了普通民众对这件事情的持续的关注，所以也就形成了举国的一个合力。”做客《今日影评》的清华大学土木学者刁基诺博士也是对这些明星艺人的果断表达了高度好评。这些霸气解约的明星艺人真是让我们感到又解气又痛快，但是当事件发酵过后，又不禁担心这些明星艺人会不会被这些品牌索要高额的违约金呢？为此，邀请了北京韬安律师事务所主任，中国演出行业协会演员经纪人委员会首席法律顾问，王军律师来为我们作出解答。\n\n“我们会有一个叫做爱国条款，那么会明确地约定，如果品牌方有违背中国的国家利益主权利益、政治立场的情形的话，艺人当然可以基于合同条款的约定提出一个单方解约权不需要承担任何赔偿责任，甚至有可能基于这个条款向品牌方追究相应的法律责任。”希望我们的艺人不会因为这些品牌的恶行而徒遭损失，也希望未来我们的艺人在代言品牌的时候一定要擦亮眼睛，而且在合同的条款上，要小心、小心、再小心！作为艺人同样也需要保护好自己的切身利益，在进行商业代言的活动的时候一定要多加小心，《今日影评》也邀请了中国广告协会广告代言人委员会秘书长张志鹏先生来提一些建议。\n\n“我觉得他们未来在去做商业代言的时候，从签约之前、签约中和签约后，都要做一些这种一些处理，比如说在签约前，其实他们要对对方的企业，包括企业的管理层，进行一个背景调查的，另外一个就是在我们签约以后，不能掉以轻心，要跟这些广告主和企业要保持一个顺畅的一个沟通。”杨幂工作室解约声明俗话讲，害人之心不可有，防人之心不可无，尤其是在品牌代言上的选择，一定要谨慎，因为明星艺人选择代言不是代言产品，而是这个品牌的理念、风格，前段时间李诞、杨笠都经历了代言翻车的时间，挑选适合自己的品牌形象去代言，对于一个明星艺人来说是非常重要的事。\n\n另一方面，我们还要做到防患于未然，不知道未来还有哪些品牌会让这些艺人无端踩雷，为此，张教授也是提出了一个解决方法。“我们是不是可以做国际品牌的红黑榜，就是说我们国家或者是说我们一个协会出面做一个相对来说比较权威的这样的一个调查，然后帮助我们的艺人规划一个比较好的范围让他们可以做一些这样的选择。”安踏推出BCI声明\n无论如何，明星艺人为了国家利益，为了民族尊严而解约，应该得到我们的支持和尊重，从长远来看，即使为之付出一些代价，也是值得的。不仅要在道义上支持明星解约，更要拿出行动，支持明星，如果企业能请这些明星代言，也是对这些辱华企业最好的回击。比如说在上周五的时候我们就看到肖战官宣了跟李宁的合作。\n\n“我觉得这是一个非常好的现象，就是第一个我们国家自有品牌民族品牌的它的力量和影响力都已经极大地提升了，第二个就是还是刚才教授说的品牌理念很关键，就是大家在合作的时候不光是一个商业行为他其实还是一个共同成长的一个过程。”明星艺人作为在社会上具有广泛关注度的公众群体，会对社会的价值观都带来较大的影响，所以时刻都要以大事为重，立场要坚定，才能树立正确的榜样。“流量就是一种社会责任，那欲戴皇冠必承其重，既然有流量，就要承担起这个社会的重量，所以我们也是期待着。我们的明星可以真正担负起自己的社会责任。”    今日观察     # 谢添意  《今日影评》当期编导国事无小事，这一次的新疆棉花事件，艺人和粉丝们的第一反应让人有些感动，在大是大非面前，国家拥有了最大的饭圈。相对之前艺人们面对突发情况的反应速度和应对方式都有了很大的进步，法律意识的提升也让他们不再成为单独的受害者，大义之前以国为先，我们同样也需要更理性的思考和预警，防范于未然。",
                "url": "http://weibo.com/ttarticle/p/show?id=2309404620387152429523"
            },
            {
                "visit_count": 12,
                "wtype": 1,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "星图数据",
                    "id": "weibo.com|3901859071",
                    "platform": "自媒体",
                    "entity": "星图数据"
                },
                "pic_urls": [
                    "https://wx2.sinaimg.cn/large/e891a4ffly4gjwxvcrg1gg203q017q3c.gif"
                ],
                "gather": {
                    "site_name": "新浪长微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069740,
                    "info_flag": [
                        "04",
                        "0408"
                    ]
                },
                "title": "星图一周资讯：工信部发布电子烟新规定；「新疆棉」事件的背后逻辑",
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 4272387247628298159,
                        "2": 698190474343604119,
                        "3": 4010731904767669704,
                        "4": 2496012457486186174,
                        "5": 4583414651141934307
                    },
                    "mentions": [
                        "syntun.com，"
                    ]
                },
                "uuid": "f5e0d5ac90fb11eba40b868712de27fd",
                "user": {
                    "uid": "3901859071",
                    "name": "星图数据"
                },
                "content": "​​欢迎关注“星图数据”！\n星图数据通过公开渠道收集、整理了2021.3.22~2021.3.28的相关行业资讯。\n\n字节跳动收购沐瞳科技扩展海外游戏市场，腾讯曾参与竞价（LinkedIn）\n3月22日，沐瞳科技CEO袁菁发表全员信，宣布沐瞳科技与字节跳动旗下游戏业务品牌朝夕光年达成战略收购协议。此次并购后，沐瞳科技会保持独立运营，将融合字节跳动在海外的运营经验，共同开拓全球游戏市场。朝夕光年相关负责人回应称，确实收购了沐瞳科技，看好这家公司的海外经验和全球化视野，沐瞳科技将继续开拓海外市场。对国内玩家来说，沐瞳科技的名字可能会稍显陌生。可旗下MOBA手游《无尽对决》，早已蜚声于东南亚，堪比国内的《王者荣耀》。有消息称腾讯也有参与并购沐瞳科技。市场猜测，腾讯此举或是抬价格，以便让字节多付成本。\n \n工信部发布电子烟新规定（LinkedIn）\n3月22日，工信部公布《关于修改<中华人民共和国烟草专卖法实施条例>的决定（征求意见稿）》，在附则中增加一条，作为第六十五条：“电子烟等新型烟草制品参照本条例中关于卷烟的有关规定执行。”工信部相关负责人表示，此次修改主要出于以下三点考虑：一是推进电子烟监管法治化；二是符合电子烟产品特性以及当前国际监管的通行做法；三是增强电子烟监管效能。\n \n四部门联合印发规定明确39种常见类型App的必要个人信息范围（证券时报）\n国家互联网信息办公室等四部门联合印发《常见类型移动互联网应用程序必要个人信息范围规定》，明确移动互联网应用程序（App）运营者不得因用户不同意收集非必要个人信息，而拒绝用户使用APP基本功能服务。《规定》明确了39种常见类型APP的必要个人信息范围，将自今年5月1日起施行。\n \n百度香港二次上市，李彦宏：百度根在中国，回家了\n3月23日百度在香港二次上市，开盘报254港元，较发行价上涨0.79%，总市值超7000亿港元。百度本次上市仪式的铜锣别出心裁，采用芯片代码锣，背板采用各式主板、转接板等集成电路板构成，其中包括百度自研芯片昆仑、鸿鹄等，涵盖了百度上市历程中的各种关键节点。\n \n百度公布本次敲锣代表为数据标注师郭梅、Apollo5G云代驾安全员雷建伟、飞桨小开发者郭佳慧三位一线业务开发者。此外还有小度机器人也同样是特殊的敲锣人。随后李彦宏进行了现场讲话，他表示，“纳斯达克只是其中一站，最终百度会回到中国来，因为我们的根在中国！今天，我们终于达成所愿，回到中国香港上市！我们回家了！”他称回到香港二次上市，是百度的再次出发，是百度的二次创业。此前百度公告，公司全球发售9500万股股份，其中香港发售股份1140万股。百度拟将全球发售募集资金净额用于，促进以人工智能为主的创新商业化，进一步发展百度移动生态，以及以支持公司的业务营运及增长。\n \n中国邮政快递报社发布报告，超五成快递员月收入不超过5000元（LinkedIn）\n中国邮政快递报社22日发布的报告显示，超五成快递员月收入不超过5000元，月收入超过1万元的仅占1.3%。超四成快递员每日派件量在100件以下，八成每日派件量不超过200件。“80后”和“90后”快递小哥是主力，数量占8成以上。51.76%的基层网点和63.86%的快递员对所属快递品牌表示“有信心”，近九成快递员对所属网点有较好的印象。\n \n市场监管总局推动乳品企业履行承诺，提升乳制品质量安全（证券时报）\n近日，市场监管总局召集中国乳制品工业协会及伊利、蒙牛、光明、飞鹤、君乐宝、三元、完达山、圣元、银桥、新希望、雀巢、美赞臣等12家大型乳企相关负责人，以视频会议形式召开乳制品质量安全提升工作会议，督促乳品企业“言必信、行必果”，进一步推进乳制品质量安全提升。\n \n胡润发布中国大消费民企百强榜，华为以1.1万亿价值位居榜首（36氪）\n据悉，胡润研究院23日发布《2021数云·胡润中国大消费民企百强榜》，华为以1.1万亿价值成为中国大消费领域价值最高民营企业。美的以5800亿元价值位列第二，海天以5000亿元价值位列第三，比亚迪以4900亿元价值排名第四，小米以4600亿元价值排名第五，蔚来、农夫山泉、格力、长城、安踏依次进入前十。\n \n快手财报：2020年营收587.8亿元，同比增长50.2%（36氪）\n据悉，快手发布2020年财报。财报显示，2020年，快手营收587.8亿元人民币，同比增长50.2%，市场预期585.82亿元；经调整亏损79.49亿元，上年同期实现经调整利润10.34亿元。快手应用的平均日活跃用户及平均月活跃用户分别为2.646亿及4.811亿，分别同比增长50.7%及45.6%。快手应用的每位日活跃用户日均使用时长由2019年的74.6分钟增加17.0%至2020年的87.3分钟。来自线上营销服务的收入由2019年的74亿元增加194.6%至219亿元；直播收入由2019年的314亿元增长5.6%至332亿元。平台上促成的电商交易的商品交易总额为3812亿元，同比增长540%；平均重复购买率由2019年的45%进一步增至2020年的65%。\n \n马斯克：新款Model S Plaid将是特斯拉至今性能最好的车辆，最多可容纳7人（智通财经）\n特斯拉CEO埃隆·马斯克表示，新款Model S Plaid将是特斯拉至今性能最好的车辆，该车型将是第一款可在2秒内加速到0-60英里每小时的量产车，还是四门配置，最多可容纳7人。然而，在当前的汽车订购选项中，该公司还没有提供7座的选项。\n \nH&M、Nike 、Adidas等“抵制”新疆棉花事件背后的逻辑：供应链主导权和定价权的争夺（LinkedIn）\n瑞典快时尚集团H&M近日在官网发表公告称，“新疆是中国最大的棉花种植区”，而H&M集团在新疆采购的棉花一直都来自经“更好的棉花计划”（BCI） 认证的农场，“由于在该地区进行可信的尽职调查变得越来越困难，BCI已决定暂停在新疆发放BCI棉花许可证”，因此该集团已不再采购新疆棉花。\n \n这则声明很快在中国社交网络引发热议。此事核心原因是新疆存在所谓“强迫劳动”、“歧视少数民族”，已由中国外交部在多个场合多次反复澄清。对此，共青团中央在官方微博点名HM，称“一边造谣抵制新疆棉花，一边又想在中国赚钱？痴心妄想”。\n \n随着事件不断发酵，京东、淘宝、天猫、唯品会都已经屏蔽了HM相关搜索。另一方面HM曾经的代言人黄轩和宋茜也相继发出声明，撇清与HM的关系。\n \n据21世纪经济报道，早在2020年9月，H&M就宣布中断与新疆华孚（A股上市公司，以纱线制造和销售为主）合作，理由是“强迫劳动”。H&M禁用新疆棉花的背后，还有Nike、Gap、Zara、UNIQLO等品牌，背后的主导者实为BCI（良好棉花发展协会）。背后是供应链主导权，以及定价权和标准的争夺。\n \nBCI是供应链联盟，掌握定价权，拥有五个类别的会员，分别是：零售品牌会员，也就是采购商，比如H&M，耐克等；供应商制造商，以棉商、纱线厂为主；种植者组织；为供应链提供技术的公司；社会团体，主要为棉花相关的非盈利性组织。BCI公布的数据显示，2019年，其零售品牌会员的棉花用量超过300万吨，占全球用量的10%，BCI的供应量则占全球的30%左右，采购量和供应量在全球首屈一指，掌控着棉花的标准和定价权。禁用新疆棉花，本质上是排斥中国供应链：中国产棉占全球约22%，新疆棉花产量占中国80%，此举目标不是新疆棉花，而是中国棉花，甚至是中国纺织供应链。服装行业分析师马岗表示，“外资品牌一边想赚中国消费者的钱，一边还打压中国供应链，低估了中国消费者的反应”。\n \n受此影响，25日安踏、李宁、特步、361度等国产品牌港股股价大涨，A股服装纺织板块走强，而耐克供应商申洲国际跌幅超5%，耐克主要合作经销商滔博一度大跌近15%。继黄轩、宋茜宣布与H&M终止合作，王一博、谭松韵与耐克终止合作后，25日，杨幂、易烊千玺、张艺兴等多位中国艺人密集宣布终止与阿迪达斯、匡威、优衣库等多家品牌一切合作。\n \n小米发布2020年业绩报告（36氪）\n据悉，小米发布2020年业绩报告。财报显示，2020年小米营收2458.66亿元，同比增长19.4%，市场预期2477.64亿元；经调整净利润130.06亿元，同比增长12.8%。2020年四季度营收704.6亿元，同比增长24.8%，经调整净利润32.04亿元，同比增长175%。\n \n此外，财报显示，小米2020年全球智能手机全年出货量为1.464亿台，同比增长17.5%。2020年智能手机收入达到人民币1522亿元，同比增长24.6%。2020年第四季度，智能手机收入达426亿元，同比增长38.4%。四季度全球智能手机出货量为4230万台，同比增长29.7%。\n \n腾讯发布年报 首次公布未成年人游戏流水占比（LinkedIn）\n3 月 24 日，腾讯在其年报中首次披露未成年人游戏流水占比：在2020年第四季度，18 岁以下未成年人对其在中国网络游戏流水占比为 6.0%，其中16岁以下未成年人的流水占比为 3.2%。2020年第四季度营收 1336.7 亿元人民币，较 2019 年第四季度同比增长 26%，高于市场预估。其中，第四季度网络游戏收入增长 29% 至人民币 391 亿元；智能手机游戏收入总额及PC端游戏收入分别为人民币367亿元及人民币102亿元。2020 年内，腾讯的总收入为人民币 4820.64 亿元，较 2019 年度增长 28%，其中网游收入则达到了 1561 亿元人民币，占比超过三成。\n\nSaaS产品方面，腾讯会议成为中国最大规模的独立云会议应用，还新推出了“腾讯会议Rooms”和“会议室连接器”，能够与客户现有的音视频设备兼容，提供高质量的互动通信体验。企业微信目前服务超过550万企业客户并与超过4亿微信用户连接。财报还显示，2020年腾讯雇员福利开支696.38亿元，去年同期531.23亿元，而员工人数从去年的62,885名增加到85,858名。\n \n新专利显示苹果汽车使用红外大灯，夜间视距增加3倍（智通财经）\n24日，苹果获得了一项夜视系统专利，该系统结合了可见光、近红外和长波红外传感器，可以全面观察前方的情况。该公司指出，限制大灯功率的法律只适用于可见光，因此红外线的大灯可以更强大。苹果表示，可以采用多种互补的图像传感技术的组合，来解决夜间或低光环境物体检测和分类的挑战。目前还不清楚苹果的汽车计划，不过如果该公司打算销售整车，那么需要准备好提供40年的支持。苹果通常会在产品停产5到7年后宣布产品过时，届时不再提供零件或维修服务。\n \n阿里巴巴：淘宝特价版已向微信提交小程序申请，正在审批中（36氪）\n3 月 25 日消息，阿里巴巴副总裁、C2M事业部总经理七公表示，阿里巴巴已提交了淘宝特价版小程序的申请，期待和腾讯的合作，但申请仍在审批中。\n \n华为全资入股持牌支付机构讯联智付，全面进军支付领域（IT之家）\n近日，据相关网站显示，深圳市讯联智付网络有限公司发生工商变更，股东上海沃芮欧信息科技有限公司退出，新增股东华为技术有限公司，持股 100%。也就是说，华为现在已经成功获得了支付牌照，顺利成为继小米之后又一家收购支付牌照的手机厂商。\n \n美团发布2020年业绩报告（36氪）\n据悉，美团发布2020年业绩报告。财报显示，2020年美团营收1147.9亿元，同比增长17.7%，市场预期1139.45亿元；经调整净利润31.2亿元，同比下降33%。2020年四季度，美团营收379.18亿元，同比增长34.7%，经调整亏损17.37亿元，同比下降163.3%。\n \n共享充电宝又集体涨价（证券时报）\n共享充电宝不约而同地又悄悄涨价了，商圈和医院等特殊场景每小时3-4元，酒吧夜店等每小时则高达10元。除此之外，封顶价格也让人咋舌，每24小时封顶价24元到40元，总封顶价99元，都赶上一个移动充电宝的价格了。\n \n苏伊士运河堵塞船只达248艘，影响厕纸咖啡等产品全球供应（澎湃新闻）\n大型货轮意外搁浅造成埃及苏伊士运河阻塞一事引发全球持续关注。截至3月27日，堵塞在运河的船只数量已经达到248艘。外媒分析认为，由于苏伊士运河的重要性，持续的堵塞将影响厕纸、咖啡等多种产品的全球供应。\n \n北京协和医院互联网医院通过北京卫健委审核，成为北京市首家获批的互联网医院（LinkedIn）\n近日，北京协和医院互联网医院通过北京市卫生健康委审核，成为北京市首家获批的互联网医院。北京协和医院互联网医院可为部分常见病、慢性病患者提供复诊服务，目前开通了心内科、内分泌科、皮肤科等19个科室，支持在院病例调阅、在线问诊、检查检验、处方开具等功能。\n \n一场可能载入美股史册的爆仓踩踏事故（LinkedIn）\n上周，一场抛售狂潮席卷美股，热门中概股大幅重挫，甚至波及多家北美媒体公司。腾讯音乐、爱奇艺、唯品会、跟谁学、百度等多只股票遭到血洗，股价巨幅震荡。到了周五，风暴进一步升级。市场最初认为，美国《外国公司问责法案》对中概股构成了利空，但有经验的投资者或许已经能感觉到这个事情不同寻常。有市场传言称，有基金通过高盛挂出多只热门股巨量卖单，出售是由“强制去杠杆”引发的。\n \n据彭博和英国《金融时报》报道，这次神秘抛售的总金额达到190亿美元，导致相关股票市值蒸发330亿美元，约合2159亿人民币。报道称，这起抛售共分为五批次完成。为了卖出这些股票，高盛及大摩的交易员一整天都在不停地打电话。高盛在盘前通过大宗交易抛售了66亿美元百度、腾讯音乐和唯品会的股票，之后在下午又卖出了ViacomCBS 、Discovery 、Farfetch Ltd.、爱奇艺和跟随学的股票，总值39亿美元。也就是说，仅高盛一家，就卖了105亿美元。其余的抛售由摩根士丹利执行，大摩当天早些时候出售了价值40亿美元的股票，下午又出售了价值40亿美元的股票。\n \n\n\n如有其他数据需求，请发送相关信息至邮箱：\ninfo@syntun.com，\n星图君会尽快回复您。\n——————————————\n星图数据是消费领域专业的大数据产品、服务和解决方案提供商；\n线上零售大数据可视化分析工具\n为品牌企业提供丰富直观的数据查询、分析和预测功能\n——————————————\n星图数据能帮您解决什么问题？\n洞悉市场变化|了解竞争对手|提升盈利能力\n——————————————\n微信号：星图数据\n英文ID：syntun​​​​",
                "url": "http://weibo.com/ttarticle/p/show?id=2309404620387190177901"
            },
            {
                "visit_count": 10,
                "wtype": 1,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "思忆悲凉动",
                    "id": "weibo.com|6189459761",
                    "platform": "自媒体",
                    "entity": "思忆悲凉动"
                },
                "pic_urls": [
                    "https://wx4.sinaimg.cn/large/006KSle9gy1gp1p90n37yj30i30k6myn.jpg",
                    "https://wx4.sinaimg.cn/large/006KSle9gy1gp1p91gzw4j30lp0avdgg.jpg",
                    "https://wx4.sinaimg.cn/large/006KSle9gy1gp1p92a5ajj30i20jl75t.jpg",
                    "https://wx4.sinaimg.cn/large/006KSle9gy1gp1p933wqgj30m10e7dgq.jpg",
                    "https://wx4.sinaimg.cn/large/006KSle9gy1gp1p93oc2lj30hx0knmyv.jpg"
                ],
                "gather": {
                    "site_name": "新浪长微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069754,
                    "info_flag": [
                        "04",
                        "0408"
                    ]
                },
                "title": "四川一李宁店顾客排长队，目击者：想买得先消费800才有资格",
                "analysis": {
                    "sentiment": -5,
                    "info_src": {},
                    "hashcode": {
                        "1": 1063296741357360771,
                        "2": 3744908797005815068,
                        "3": 4567605056791965325,
                        "4": 7930847160410732215,
                        "5": 6448189355201419272
                    }
                },
                "uuid": "fe79adec90fb11ebba30ae475236662c",
                "user": {
                    "uid": "6189459761",
                    "name": "思忆悲凉动"
                },
                "content": "​​最近，因为“HM碰瓷新疆棉花”惹怒国人，许多地区实体店铺面临关门凉凉的局面。深挖背后，还有阿迪、耐克等，一时间网友愤怒不已，对于这群外国品牌“吃中国饭砸中国碗”的行为，大家感到鄙视，并且表示今后要加大力度支持国产品牌。\n所以这段时间可以看到，某些国产品牌崛起，甚至出现连夜排队购买的画面。\n\n3月27日，四川成都，一位拍摄者表示某地步行街一家李宁门店外排起长队，不少人头一天晚上就开始来排。\n而旁边的阿迪却门可罗雀，店铺内的顾客寥寥无几，画风形成了鲜明的对比，由衷感慨国人的团结一心。\n此外，拍摄者无奈称，排队进李宁，人均消费2000左右，想要买限量版，必须先买一双800多元的鞋子。\n\n拿2000元买鞋，对普通工薪阶层的上班族来说有些奢侈了。虽然李宁整体要比阿迪便宜很多，但却架不住“限量版”的诱惑，许多顾客都是冲着限量版来买的。\n限量版发售的多了就不是“限量款”了，顾客买东西一定要理智。\n“捧得太高了不是啥好事，大家正常购买就好了，否则回头该涨价了。”\n评论中有一位网友说出了大实话，商人的本质是赚钱，当市场商品供不应求的时候，资本很有可能会抬高价格，涨价只是时间早晚而已。\n\n很多人去购买都是源于一腔爱国心，购买力强说明国民经济实力增强，但客观的来讲，对任何一个品牌都不能过分“捧”，热情要有一个度，过了头可能会有相反的效果。\n大家对国产品牌热情高涨是件好事，带动了经济，让他们日益强大，如果只是某几家品牌独大，担心日后可能会出现价格上涨等情况，到那时辜负了顾客的期待就不好了。所以，其他国产品牌也要加把劲，争取出现百家齐放的画面，让顾客们“雨露均沾”。\n\n时代发达的今天，无论是鞋衣服还是包，质量都大同小异，不会有多么明显的质量差距，关键是看品牌效应。\n一个品牌想要获得好的口碑不能图赚快钱，想要走得长远必须深得民心，无论从价格还是质量都要讲良心。\n\n经过了这次HM事件，对国产品牌来说是一次好的契机，将国产品牌发扬壮大，我们的产品质量一直都很好，只是缺少一个机会罢了，这次事件带来的影响力，希望能让国产品牌更一次深入民心。\n同时，我们希望国产品牌不要骄傲，不能因为顾客的支持就提高市场价格，不奢求一飞冲天，只希望能稳稳上升，不辜负国人的期待。\n​\n本文部分内容转载于网络，如有侵权，麻烦联系删除！​​​​",
                "url": "http://weibo.com/ttarticle/p/show?id=2309404620387207217425"
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387142338174",
                "pid": "4618389982546373",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069604,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 2987203498073623101,
                        "2": 1645103905065330560,
                        "3": 4536560375394389048,
                        "4": 7407775761963770821,
                        "5": 1202187872022564531
                    },
                    "mentions": [
                        "李现ing"
                    ],
                    "hashtag": [
                        "我支持新疆棉花"
                    ]
                },
                "reply_count": 0,
                "uuid": "a45587d290fb11ebad8206af6926cef9",
                "content": "11//@李现ing:#我支持新疆棉花#",
                "url": "http://weibo.com/1884676497/K8ymu9Ogu",
                "retweeted": {
                    "repost_count": 37465585,
                    "wtype": 1,
                    "like_count": 0,
                    "mid": "4618389982546373",
                    "ctime": 1616593440,
                    "pic_urls": [
                        "http://wx2.sinaimg.cn/thumbnail/0033ImPzly1govbvfxn40j60u01hcn7u02.jpg"
                    ],
                    "reply_count": 145341,
                    "user": {
                        "friends_count": 3056,
                        "uid": 2803301701,
                        "gender": "m",
                        "level": 48,
                        "bi_followers_count": 0,
                        "statuses_count": 130418,
                        "followers_count": 129085181,
                        "verified": 2,
                        "name": "人民日报",
                        "location": [
                            "北京"
                        ],
                        "verified_type": 3
                    },
                    "content": "[话筒]#我支持新疆棉花# ​​",
                    "url": "http://weibo.com/2803301701/K7IpgaGqx"
                },
                "wtype": 2,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "刘向宁小窝",
                    "id": "weibo.com|1884676497",
                    "platform": "自媒体",
                    "entity": "刘向宁小窝"
                },
                "device": "荣耀30 5G",
                "user": {
                    "friends_count": 348,
                    "profile_img_url": "http://tva3.sinaimg.cn/crop.0.0.180.180.50/7055e191jw8eye3lc16u1j20500503yn.jpg",
                    "gender": "m",
                    "level": 3,
                    "verified": 0,
                    "created_at": 1291637015,
                    "verified_type": -1,
                    "uid": 1884676497,
                    "lang_code": "zh-cn",
                    "location_code": "93415",
                    "bi_followers_count": 4,
                    "statuses_count": 99,
                    "followers_count": 597,
                    "name": "刘向宁小窝",
                    "location": [
                        "辽宁省",
                        "盘锦市"
                    ]
                }
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387147319388",
                "pid": "4618389982546373",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069604,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 4009591967248171543,
                        "2": 1117074761443878477,
                        "3": 1482030087110493394,
                        "4": 7407775761963770821,
                        "5": -7407775761963770821
                    },
                    "mentions": [
                        "杨幂",
                        "杨幂把我抱在怀里"
                    ],
                    "hashtag": [
                        "我支持新疆棉花"
                    ]
                },
                "reply_count": 0,
                "uuid": "a4b8ee6290fb11eb8f92f6f1094b8d63",
                "content": "衣带渐宽终不悔，为伊消得人憔悴。//@杨幂把我抱在怀里:#我支持新疆棉花# [心]//@杨幂:#我支持新疆棉花# ​​​",
                "url": "http://weibo.com/7129098906/K8ymuuI6E",
                "retweeted": {
                    "repost_count": 37465592,
                    "wtype": 1,
                    "like_count": 0,
                    "mid": "4618389982546373",
                    "ctime": 1616593440,
                    "pic_urls": [
                        "http://wx2.sinaimg.cn/thumbnail/0033ImPzly1govbvfxn40j60u01hcn7u02.jpg"
                    ],
                    "reply_count": 145341,
                    "user": {
                        "friends_count": 3056,
                        "uid": 2803301701,
                        "gender": "m",
                        "level": 48,
                        "bi_followers_count": 0,
                        "statuses_count": 130418,
                        "followers_count": 129085181,
                        "verified": 2,
                        "name": "人民日报",
                        "location": [
                            "北京"
                        ],
                        "verified_type": 3
                    },
                    "content": "[话筒]#我支持新疆棉花# ​​",
                    "url": "http://weibo.com/2803301701/K7IpgaGqx"
                },
                "wtype": 2,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "聪明_铅笔",
                    "id": "weibo.com|7129098906",
                    "platform": "自媒体",
                    "entity": "聪明_铅笔"
                },
                "device": "Android客户端",
                "user": {
                    "friends_count": 71,
                    "profile_img_url": "http://tvax1.sinaimg.cn/crop.0.0.512.512.50/007MsYiely8gngj1x7fzhj30e80e8go7.jpg",
                    "gender": "f",
                    "level": 7,
                    "verified": 0,
                    "created_at": 1557474023,
                    "verified_type": -1,
                    "uid": 7129098906,
                    "lang_code": "zh-cn",
                    "location_code": "8447",
                    "bi_followers_count": 36,
                    "statuses_count": 84051,
                    "followers_count": 599,
                    "name": "聪明_铅笔",
                    "location": [
                        "河北省"
                    ]
                }
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387147055732",
                "pid": "4618389982546373",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069603,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 510148138210074375,
                        "2": 3768409448567570343,
                        "3": 2255788719113008893,
                        "4": 7407775761963770821,
                        "5": -7407775761963770821
                    },
                    "mentions": [
                        "深情的高伟光",
                        "深情的高伟光♥️#高伟光光荣与梦想#♥️#高伟光#♥️#高伟光王居安#♥️#高伟光三生三世枕上书#♥️#高伟光东华帝君#",
                        "追光耗子-satine"
                    ],
                    "hashtag": [
                        "高伟光中国科普月环保大使",
                        "我支持新疆棉花",
                        "高伟光",
                        "高伟光三生三世枕上书",
                        "高伟光东华帝君",
                        "高伟光益起来",
                        "高伟光王居安",
                        "高伟光光荣与梦想"
                    ]
                },
                "reply_count": 0,
                "uuid": "a3af426490fb11eb893a0222092cc5d8",
                "content": "高伟光正能量@深情的高伟光♥️#高伟光光荣与梦想#♥️#高伟光#♥️#高伟光王居安#♥️#高伟光三生三世枕上书#♥️#高伟光东华帝君# //@追光耗子-satine:#高伟光益起来#支持新疆棉花@深情的高伟光 #高伟光中国科普月环保大使# //@深情的高伟光:#我支持新疆棉花# [心]",
                "url": "http://weibo.com/6287753640/K8ymutBw8",
                "retweeted": {
                    "repost_count": 37465584,
                    "wtype": 1,
                    "like_count": 0,
                    "mid": "4618389982546373",
                    "ctime": 1616593440,
                    "pic_urls": [
                        "http://wx2.sinaimg.cn/thumbnail/0033ImPzly1govbvfxn40j60u01hcn7u02.jpg"
                    ],
                    "reply_count": 145341,
                    "user": {
                        "friends_count": 3056,
                        "uid": 2803301701,
                        "gender": "m",
                        "level": 48,
                        "bi_followers_count": 0,
                        "statuses_count": 130418,
                        "followers_count": 129085181,
                        "verified": 2,
                        "name": "人民日报",
                        "location": [
                            "北京"
                        ],
                        "verified_type": 3
                    },
                    "content": "[话筒]#我支持新疆棉花# ​​",
                    "url": "http://weibo.com/2803301701/K7IpgaGqx"
                },
                "wtype": 2,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "深情的耗子t-satine",
                    "id": "weibo.com|6287753640",
                    "platform": "自媒体",
                    "entity": "深情的耗子t-satine"
                },
                "device": "Redmi 9A",
                "user": {
                    "friends_count": 334,
                    "profile_img_url": "http://tvax1.sinaimg.cn/crop.0.0.1080.1080.50/006RwLXily8gl6x1349l0j30u00u0myu.jpg",
                    "gender": "m",
                    "level": 7,
                    "verified": 0,
                    "created_at": 1497755694,
                    "verified_type": -1,
                    "uid": 6287753640,
                    "lang_code": "zh-cn",
                    "bi_followers_count": 57,
                    "statuses_count": 38884,
                    "followers_count": 517,
                    "name": "深情的耗子t-satine"
                }
            },
            {
                "repost_count": 0,
                "like_count": 0,
                "mid": "4620387146793631",
                "pid": "4618389982546373",
                "gather": {
                    "site_name": "新浪微博",
                    "site_domain": "weibo.com",
                    "stime": 1617069602,
                    "info_flag": [
                        "04",
                        "0401"
                    ]
                },
                "analysis": {
                    "sentiment": 0,
                    "info_src": {},
                    "hashcode": {
                        "1": 2458215849234575072,
                        "2": 4081162987012298424,
                        "3": 732435040738600256,
                        "4": 7407775761963770821,
                        "5": -7407775761963770821
                    },
                    "mentions": [
                        "念念念雨伞",
                        "THE9-刘雨昕",
                        "刘雨昕粉丝后援会"
                    ],
                    "hashtag": [
                        "我支持新疆棉花"
                    ]
                },
                "reply_count": 0,
                "uuid": "a3d4932090fb11eb98709ad5d42ead84",
                "content": "[吃瓜]//@念念念雨伞:支持//#我支持新疆棉花# //@刘雨昕粉丝后援会:#我支持新疆棉花# ​​​//@THE9-刘雨昕:#我支持新疆棉花# ​​​",
                "url": "http://weibo.com/6541415272/K8ymusvkH",
                "retweeted": {
                    "repost_count": 37465586,
                    "wtype": 1,
                    "like_count": 0,
                    "mid": "4618389982546373",
                    "ctime": 1616593440,
                    "pic_urls": [
                        "http://wx2.sinaimg.cn/thumbnail/0033ImPzly1govbvfxn40j60u01hcn7u02.jpg"
                    ],
                    "reply_count": 145341,
                    "user": {
                        "friends_count": 3056,
                        "uid": 2803301701,
                        "gender": "m",
                        "level": 48,
                        "bi_followers_count": 0,
                        "statuses_count": 130418,
                        "followers_count": 129085181,
                        "verified": 2,
                        "name": "人民日报",
                        "location": [
                            "北京"
                        ],
                        "verified_type": 3
                    },
                    "content": "[话筒]#我支持新疆棉花# ​​",
                    "url": "http://weibo.com/2803301701/K7IpgaGqx"
                },
                "wtype": 2,
                "ctime": 1617069600,
                "publisher": {
                    "site_name": "微博",
                    "name": "树下花儿不开馨VL97",
                    "id": "weibo.com|6541415272",
                    "platform": "自媒体",
                    "entity": "树下花儿不开馨VL97"
                },
                "device": "三星android智能手机",
                "user": {
                    "friends_count": 40,
                    "profile_img_url": "http://tvax1.sinaimg.cn/default/images/default_avatar_female_50.gif",
                    "gender": "f",
                    "level": 4,
                    "verified": 0,
                    "created_at": 1525407885,
                    "verified_type": -1,
                    "uid": 6541415272,
                    "lang_code": "zh-cn",
                    "bi_followers_count": 0,
                    "statuses_count": 10629,
                    "followers_count": 505,
                    "name": "树下花儿不开馨VL97",
                    "location": [
                        "海外"
                    ]
                }
            }
        ]
    },
    "success": true,
    "message": "success"
}`)
	var resp Resp

	err := json.Unmarshal(data, &resp)
	if err !=
		
	

	for _, item :=              range r e s p . D a t a . L i s t {
		fmt.Print(item.Gather.SiteName)
		// switch {
		// 	case item.Gather. == "微信":
		// 		WBPersing(item)
		// 	case item.Gather.Gather == "新浪微博":
		// 		WXPersing(item)
		// 	default:
		// 		APPPersing(item)
		// }
	}
	// fmt.Printf("%#v\n", resp)
}














