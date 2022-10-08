package main

import (
	"fmt"

	"xorm.io/builder"
)

func main() {
	sql, args, _ := builder.ToSQL(builder.Exists(builder.Select("a").From("table")))
	fmt.Println(sql, args)
	sql, args, _ = builder.ToSQL(builder.NotExists(builder.Select("a").From("table")))
	fmt.Println(sql, args)

	var targetGroupTypeSql string
	var s = []string{
		"1231",
		"ksdfjls",
	}

	for i, targetGroupType := range s {
		targetGroupTypeSql = fmt.Sprintf("%s JSON_CONTAINS(expression_json, JSON_OBJECT('type','%s'))", targetGroupTypeSql, targetGroupType)
		if i < (len(s) - 1) {
			targetGroupTypeSql = fmt.Sprintf("%s or", targetGroupTypeSql)
		}
	}
	fmt.Println(targetGroupTypeSql)
}
