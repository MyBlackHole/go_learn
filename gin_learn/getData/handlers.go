package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	var postjson Json
	if err := c.ShouldBindJSON(&postjson); err != nil {
		c.JSON(http.StatusBadRequest, MyResp{
			ERROR: err.Error(),
		})
		return
	}

	// 写redis日志
	go save(postjson, 60*60*24*7*time.Second)

	var resp Resp
	if postjson.Status {
		var person *Person
		person = &Person{
			Json: postjson,
			Token: Token{
				Token: "4b7c350e448945fa8ab9c8ed22110a80",
			},
		}

		err := checkPerson(person)
		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}

		resp, err = request(person)
		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}
	} else {
		var xgperson XGPerson

		err := jsonToXGPerson(&postjson, &xgperson)

		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}

		err = checkXGPerson(&xgperson)
		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}

		xgresp, err := xgRequest(&xgperson)
		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}

		err = XGRespToResp(&resp, &xgresp)
		if err != nil {
			c.JSON(http.StatusBadRequest, MyResp{
				ERROR: err.Error(),
			})
			return
		}
	}

	var myresp MyResp
	mainPersing(&resp, &myresp, postjson.Original)

	// 发包到底层
	if postjson.SendStatus {
		go package_main(myresp)
	}

	c.JSON(http.StatusOK, myresp)
}

func getRedisData(c *gin.Context) {
	DefaultKey := CurrentDate()
	key := c.DefaultQuery("key", DefaultKey)
	resp, err := redisdb.SMembers(key).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, RedisResp{
			ERROR: err.Error(),
		})
		return
	}

	var redisResp RedisResp

	var List []RedisData
	for _, redisDataStr := range resp {
		var redisData RedisData
		err = json.Unmarshal([]byte(redisDataStr), &redisData)
		if err != nil {
			redisResp.ERROR = redisResp.ERROR + "|" + err.Error()
		}
		List = append(List, redisData)
	}
	redisResp.Data = List

	c.JSON(http.StatusOK, redisResp)
}
