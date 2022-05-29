package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 空json
func readUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// 保存数据
func saveData(c *gin.Context) {
	var result Result
	if err := c.ShouldBindJSON(&result); err != nil {
		log.Info("saveData.ShouldBindJSON" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Info("saveData.Marshal" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	key := getMd5(result.KeyWord)
	err = redisdb.Set(key+result.CrawlerType, string(data), -1*time.Second).Err()
	if err != nil {
		log.Info("saveData.redisdb.Set" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
	})
}

// 获取数据
func getData(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Info("getData.ShouldBindJSON" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	key := getMd5(person.KeyWord)
	data, err := redisdb.Get(key + person.CrawlerType).Result()
	if err != nil {
		log.Info("getData.redisdb.Get" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	dataByt := []byte(data)
	var result Result
	err = json.Unmarshal(dataByt, &result)
	if err != nil {
		log.Info("getData.Unmarshal" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// 添加日志记录
func addChargeLog(c *gin.Context) {
    var chargeLog map[string]interface{}
	chargeLog = make(map[string]interface{})
	
	if err := c.Bind(&chargeLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	// 添加时间
	chargeLog["Time"] = CurrentDateTime()

	collectionNmae, ok := chargeLog["collectionName"].(string)
	if ok == false {
		collectionNmae = "toll_log"
	}
	delete(chargeLog, "collectionName")

	collection := Databases.C(collectionNmae)
	err := collection.Insert(chargeLog)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"success": true,
	})
}

// 查询日志记录
func findChargeLog(c *gin.Context) {
    var param map[string]interface{}
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	collectionNmae, ok := param["collectionName"].(string)
	if ok == false {
		collectionNmae = "toll_log"
	}

	// 排序操作
	// Sort, Sort_ok := param["Sort"].(string)
	SortList, Sort_ok:= param["Sort"].([]interface{})

	// 个数
	Limit, Limit_ok := param["Limit"].(float64)

	delete(param, "collectionName")
	delete(param, "Sort")
	delete(param, "Limit")

	collection := Databases.C(collectionNmae)
    var all []map[string]interface{}
    // err := collection.Find(param).All(&all)

    query := collection.Find(param)
	if Sort_ok {
		for _, Sort :=range SortList {
			query = query.Sort(Sort.(string))
		}
	}

	if Limit_ok {
		query = query.Limit(int(Limit))
	}

	err := query.All(&all)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, all)
}
