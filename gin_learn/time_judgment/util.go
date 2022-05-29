package main

import (
	"encoding/hex"
	"crypto/md5"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func JsonToMap(jsonStr string) (map[string]string, error) {
    m := make(map[string]string)
    err := json.Unmarshal([]byte(jsonStr), &m)
    if err != nil {
        log.Printf("Unmarshal with error: %+v\n", err)
        return nil, err
    }
 
    for k, v := range m {
        log.Printf("%v: %v\n", k, v)
    }
 
    return m, nil
}

// 获取 string md5
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
