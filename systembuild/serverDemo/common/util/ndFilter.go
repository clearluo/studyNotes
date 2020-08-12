package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"serverDemo/common/log"
	"time"
)

const (
	id      = "0"
	ct      = "qmbz"
	key     = "c114d1513731753d3d8ed0a807ae2126"
	baseUrl = "https://chat.99.com/chat?action=filter_public_list"
)

var (
	dictDir   = filepath.Join("config", "sensitive")
	FilterStr = ""
)

func GetFilterList() error {
	ts := fmt.Sprintf("%v", time.Now().Unix())
	sig := EncodeMd5(id + ct + ts + key)
	url := baseUrl + "&id=" + id
	url += "&ct=" + ct
	url += "&ts=" + ts
	url += "&sig=" + sig
	data, err := GetHtml(url)
	if len(data) < 1000 || err != nil {
		err := fmt.Errorf("get err:", url)
		log.Warn(err)
		return err
	}
	FilterStr = string(data)
	if !PathExists(dictDir) {
		if err := os.MkdirAll(dictDir, os.ModePerm); err != nil {
			err := fmt.Errorf("dir not found，please mkdir dir:%v", dictDir)
			log.Error(err)
			return err
		}
	}
	if err := ioutil.WriteFile(filepath.Join(dictDir, "default.txt"), []byte(data), 0666); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func ReadFlterFromFile() error {
	filePath := filepath.Join(dictDir, "default.txt")
	data, err := ioutil.ReadFile(filePath)
	if len(data) < 1000 || err != nil {
		err := fmt.Errorf("readFile err:", filePath)
		log.Warn(err)
		panic(err)
		return err
	}
	FilterStr = string(data)
	return nil
}
