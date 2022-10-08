package main

import (
	"encoding/json"
)

// 需要我们去重写ToDB、FromDB方法
type Cover struct {
	Id   int64  `json:"id,omitempty"`
	Fid  string `json:"fid,omitempty"`
	Type int8   `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

func (c *Cover) FromDB(bytes []byte) error {
	return json.Unmarshal(bytes, c)
}

func (c *Cover) ToDB() (bytes []byte, err error) {
	bytes, err = json.Marshal(c)
	return
}

// type Course struct {
// 	Id            int64             `json:"id,string" form:"id"`
// 	Name          string            `json:"name" form:"name"`
// 	Brief         string            `json:"brief" form:"brief"`
// 	Description   string            `json:"description" form:"description"`
// 	Cover         common.Cover      `xorm:"Text" json:"cover" form:"cover"`
// 	Categories    common.Int64Array `xorm:"Text" json:"categories" form:"categories[]"`
// 	Tags          common.Int64Array `json:"tags" form:"tags[]"`
// 	Difficulty    float64           `json:"difficulty" form:"difficulty"`
// 	Price         float64           `json:"price" form:"price"`
// 	Markets       common.Int64Array `json:"markets" form:"markets[]"`
// 	StudentAmount int64             `json:"studentAmount" form:"studentAmount"`
// 	SubjectAmount int64             `json:"subjectAmount" form:"subjectAmount"`
// 	Crt           time.Time         `json:"crt"`
// 	Lut           time.Time         `json:"lut"`
// 	Status        int16             `json:"status" form:"status"`
// 	common.Page   `xorm:"-"`
// }
