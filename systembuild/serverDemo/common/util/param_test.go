package util

import (
	"testing"
)

func TestUtil(t *testing.T) {
	data := map[string]interface{}{
		"cpu": 100,
		"mem": 20,
	}
	obj := map[string]interface{}{
		"data": data,
	}
	param, err := GetValue("data.cpu", obj)
	t.Log("param", err, param)
}
