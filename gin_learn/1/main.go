package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

type Result struct {
	Keyword   string   `json:"keyword"`
	Token     string   `json:"token"`
	StartTime string   `json:"startTime"`
	EndTime   string   `json:"endTime"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
	Uids      []string `json:"uids"`
	SiteTypes []string `json:"siteTypes"`
}

type Person struct {
	Keyword   string   `json:"keyword"`
	StartTime string   `json:"startTime"`
	EndTime   string   `json:"endTime"`
	SiteTypes []string `json:"siteTypes"`
}

func main() {
	r := gin.Default()
	r.POST("/get_total", getTotal)
	// r.GET("/", )
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getTotal(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	type result Result
	_json := &result{
		Keyword:   person.Keyword,
		Token:     "99999999999999999999999999999999",
		StartTime: person.StartTime,
		EndTime:   person.EndTime,
		PageIndex: 1,
		PageSize:  1,
		Uids:      []string{},
		SiteTypes: person.SiteTypes,
	}
	data, _ := json.Marshal(_json)

	url := "http://0000000000000:8081/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var total int = 0

	totalresp, err := gojsonq.New().FromString(string(body)).FindR("data.total")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = totalresp.As(&total)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"total":   total,
	})
}
