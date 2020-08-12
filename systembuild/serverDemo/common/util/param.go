package util

import (
	"strings"
)

func GetValues(params []string, obj map[string]interface{}) (res []interface{}, flag bool) {
	for _, key := range params {
		item, _ := GetValue(key, obj)
		//if !flag {
		//	return res, false
		//}
		res = append(res, item)
	}
	return
}

func GetValue(param string, obj map[string]interface{}) (res interface{}, flag bool) {
	params := strings.Split(param, ".")
	if len(params) < 2 {
		if val, err := obj[params[0]]; err {
			return val, true
		}
	}
	for _, key := range params {
		res, flag = obj[key]
		if !flag {
			return nil, false
		}
		if temp, f := res.(map[string]interface{}); f {
			obj = temp
		}
	}
	return
}
