package main

import (
	"strconv"
)

func interface2String(inter interface{}) (s string) {
	switch inter.(type) {
	case string:
		s = inter.(string)
		break
	case int:
		s = strconv.Itoa(inter.(int))
		break
	case float64:
		s = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	default:
		s = ""
	}
	return
}
