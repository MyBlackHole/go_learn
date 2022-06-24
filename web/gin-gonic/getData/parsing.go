package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func mainPersing(resp *Resp, myresp *MyResp, original bool) {
	myresp.Total = resp.Data.Total
	var List []URunBlog
	for _, blog := range resp.Data.List {
		if blog.Gather.SiteName == "微信" {
			var urunBlog URunBlog
			WXPersing(&blog, &urunBlog)
			List = append(List, urunBlog)
		} else if blog.Gather.SiteDomain == "weibo.com" {
			// 原文解析
			var urunBlog URunBlog
			WBPersing(&blog, &urunBlog, false)

			// 给转发拼接原文content
			var strSuffix string

			// orc解析
			if blog.OCR != "" {
				var ocrUrunBlog URunBlog
				WBPersing(&blog, &ocrUrunBlog, true)
				if !original {
					List = append(List, ocrUrunBlog)
				}
				strSuffix += ocrUrunBlog.Content
			}

			// 转发解析
			var quoteUrunBlog URunBlog
			if blog.Retweeted != nil {
				WBPersing(blog.Retweeted, &quoteUrunBlog, false)
				if !original {
					List = append(List, quoteUrunBlog)
				}
				strSuffix += quoteUrunBlog.Content
				mergeQuote(&urunBlog, &quoteUrunBlog)
			}

			setWeiboType(&urunBlog)

			List = append(List, urunBlog)
		} else {
			var urunBlog URunBlog
			APPPersing(&blog, &urunBlog)
			List = append(List, urunBlog)
		}
	}
	myresp.List = List
}

func setWeiboType(urunBlog *URunBlog) {
	if urunBlog.Quoteauthor == "" {
		urunBlog.Weibotype += "1"
	} else {
		urunBlog.Weibotype += "2"
	}

	if urunBlog.Imageurl != "" || urunBlog.Quoteimageurl != "" {
		urunBlog.Weibotype += " 5"
	}

	if urunBlog.Videourl != "" {
		urunBlog.Weibotype += " 6"
	}

	if urunBlog.Imageurl == "" && urunBlog.Quoteimageurl == "" && urunBlog.Videourl == "" {
		urunBlog.Weibotype += " 4"
	}
}

func mergeQuote(urunBlog *URunBlog, quoteUrunBlog *URunBlog) {
	urunBlog.Quoteauthor = quoteUrunBlog.Name
	urunBlog.Quoteblogid = quoteUrunBlog.Blogid
	urunBlog.Quotecontent = quoteUrunBlog.Content
	urunBlog.Quoteuid = quoteUrunBlog.UID
	urunBlog.Quotetime = quoteUrunBlog.Time
	urunBlog.Quoteurl = quoteUrunBlog.URL
	urunBlog.Quotetransmits = quoteUrunBlog.Transmits
	urunBlog.Quotecomments = quoteUrunBlog.Comments
	urunBlog.Quotepraises = quoteUrunBlog.Praises
	urunBlog.Quoteportraiturl = quoteUrunBlog.Portraiturl
	urunBlog.Quoteimageurl = quoteUrunBlog.Imageurl

	urunBlog.Title = urunBlog.Title + "\n" + quoteUrunBlog.Title

	// 2021-06-15 停止拼接原文与转发
	// urunBlog.Content = urunBlog.Content + "\u2225" + quoteUrunBlog.Content
}

func WBPersing(blog *Blog, urunBlog *URunBlog, ocrStatus bool) {
	urunBlog.Name = blog.User.Name
	urunBlog.Hashcode = getMd5(blog.Content)
	urunBlog.UID = interface2String(blog.User.UID)
	urunBlog.Addon = getTimeNow13()
	urunBlog.Time = int64(blog.Ctime * 1000.0)
	urunBlog.Transmits = blog.RepostCount
	urunBlog.Comments = blog.ReplyCount
	urunBlog.Praises = blog.LikeCount
	urunBlog.Portraiturl = blog.User.ProfileImgURL
	urunBlog.Imageurl = strings.Join(blog.PicUrls, "|")
	urunBlog.Sitetype = 7

	urunBlog.Title = blog.Content

	i, err := strconv.Atoi(blog.Mid)
	if err != nil {
		i = 0000000000000000
		fmt.Printf("mid-[%s]\n", blog.Mid)
	}
	urunBlog.Blogid = MidToBid(i)

	urunBlog.URL = fmt.Sprintf("http://weibo.com/%s/%s", urunBlog.UID, urunBlog.Blogid)

	// 判别weibo文章
	if isttarticle := strings.Index(blog.URL, "ttarticle"); isttarticle != -1 {
		r, _ := regexp.Compile("([0-9]*$)")
		id := r.FindString(blog.URL)
		urunBlog.URL = blog.URL
		urunBlog.Blogid = id
	}

	// 判别ocr
	if ocrStatus {
		var ocrPrefix string = "ocr"
		urunBlog.ID = ocrPrefix + getMd5(urunBlog.URL)
		urunBlog.Content = ocrPrefix + blog.OCR
	} else {
		urunBlog.Content = blog.Content
		urunBlog.ID = getMd5(urunBlog.URL)
	}

	urunBlog.Taskid = 999999999
	count := len(blog.User.Location)
	switch {
	case count == 1:
		urunBlog.Province = blog.User.Location[0]
	case count == 2:
		urunBlog.Province = blog.User.Location[0]
		urunBlog.City = blog.User.Location[1]
	default:
		urunBlog.Province = ""
	}


	urunBlog.Sitetype = 7
	urunBlog.Language = 1
	urunBlog.From = "weibo.com"
}

func WXPersing(blog *Blog, urunBlog *URunBlog) {
	urunBlog.Name = blog.User.Name
	urunBlog.Hashcode = getMd5(blog.Title)
	urunBlog.Account = interface2String(blog.User.UID)
	urunBlog.Addon = getTimeNow13()
	urunBlog.Time = int64(blog.Ctime * 1000)
	urunBlog.URL = blog.URL
	urunBlog.Title = blog.Title
	urunBlog.Content = blog.Content
	urunBlog.Comments = blog.ReplyCount
	urunBlog.Taskname = "微信_" + blog.User.Name
	urunBlog.Imageurl = strings.Join(blog.PicUrls, "|")
	urunBlog.ID = getMd5(blog.URL)
	urunBlog.Domain = blog.Gather.SiteDomain

	urunBlog.Taskid = 888888888
	count := len(blog.User.Location)
	switch {
	case count == 1:
		urunBlog.Province = blog.User.Location[0]
	case count == 2:
		urunBlog.Province = blog.User.Location[0]
		urunBlog.City = blog.User.Location[1]
	default:
		urunBlog.Province = ""
	}

	urunBlog.Groupname = "微信"
	urunBlog.Sitetype = 14
	urunBlog.Language = 1
}

func APPPersing(blog *Blog, urunBlog *URunBlog) {
	urunBlog.Hashcode = getMd5(blog.Title)
	urunBlog.Account = interface2String(blog.User.UID)
	urunBlog.Addon = getTimeNow13()
	urunBlog.Time = int64(blog.Ctime * 1000)
	urunBlog.URL = blog.URL
	urunBlog.Title = blog.Title
	urunBlog.Content = blog.Content
	urunBlog.Comments = blog.ReplyCount
	urunBlog.Taskname = blog.User.Name
	urunBlog.Imageurl = strings.Join(blog.PicUrls, "|")
	urunBlog.ID = getMd5(blog.URL)
	urunBlog.Domain = blog.Gather.SiteDomain

	urunBlog.Taskid = 777777777
	count := len(blog.User.Location)
	switch {
	case count == 1:
		urunBlog.Province = blog.User.Location[0]
	case count == 2:
		urunBlog.Province = blog.User.Location[0]
		urunBlog.City = blog.User.Location[1]
	default:
		urunBlog.Province = ""
	}

	urunBlog.Groupname = blog.User.Name
	urunBlog.Sitetype = 12
	urunBlog.Language = 1
}
