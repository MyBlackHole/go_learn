package main

type Result struct {
	KeyWord     string `json:"KeyWord"`
	CrawlerType string `json:"CrawlerType"`
	LastTime    string `json:"LastTime"`
	BlogID      string `json:"BlogID"`
}

type Person struct {
	KeyWord     string `json:"KeyWord"`
	CrawlerType string `json:"CrawlerType"`
}

// type ChargeLog struct {
// 	Request interface{} `json:"Request"`
// 	Resp    interface{} `json:"Resp"`
// 	Time    string      `json:"Time"`
// }
