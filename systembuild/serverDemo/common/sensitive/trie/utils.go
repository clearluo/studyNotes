package trie

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"serverDemo/common/log"
	"serverDemo/common/util"
	"strings"
	"time"
)

/**
 * 参考文档：https://github.com/huayuego/wordfilter
 */
var GblackTrie *Trie
var whitePrefixTrie *Trie
var whiteSuffixTrie *Trie
var dictDir string = filepath.Join("config/sensitive/")

// InitAllTrie 初始化三种Trie
func InitTrie() {
	if err := util.GetFilterList(); err != nil {
		util.ReadFlterFromFile()
	}
	BlackTrie() // monitorSrv服需要手动拷贝敏感词库，不能从项目excel中读取
	wordSli := GblackTrie.ReadAll()
	if len(wordSli) < 100 {
		err := fmt.Errorf("sensitive length too short:%d", len(wordSli))
		fmt.Println(err)
		log.Error(err)
		panic(err)
	}
	//WhitePrefixTrie()
	//WhiteSuffixTrie()
}

// BlackTrie 返回黑名单Trie树
func BlackTrie() *Trie {
	fileName := filepath.Join(dictDir, "default.txt")
	_, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		panic(err)
	}
	//wordSli:=[]string{}
	//if err:=json.Unmarshal(contentByte,&wordSli);err!=nil{
	//	fmt.Println(err,"BBBBBBBBBBBBBBBBB")
	//	log.Error(err)
	//	panic(err)
	//}
	//if len(wordSli)<100{
	//	err:=fmt.Errorf("sensitive length too short:",len(wordSli))
	//	fmt.Println(err)
	//	log.Error(err)
	//	panic(err)
	//}
	if GblackTrie == nil {
		GblackTrie = NewTrie()
		GblackTrie.CheckWhiteList = true

		loadDict(GblackTrie, "add", dictDir)
		//loadDict(blackTrie, "del", "config/sensitive/black/exclude")
	}
	return GblackTrie
}

// WhitePrefixTrie 返回白名单前缀Trie树
func WhitePrefixTrie() *Trie {
	if whitePrefixTrie == nil {
		whitePrefixTrie = NewTrie()
		//loadDict(whitePrefixTrie, "add", "config/sensitive/white/prefix")
		loadDict(whitePrefixTrie, "add", "config/sensitive")
	}
	return whitePrefixTrie
}

// ClearWhitePrefixTrie 清空白名单前缀Trie树
func ClearWhitePrefixTrie() {
	whitePrefixTrie = NewTrie()
}

// WhiteSuffixTrie 返回白名单后缀Trie树
func WhiteSuffixTrie() *Trie {
	if whiteSuffixTrie == nil {
		whiteSuffixTrie = NewTrie()
		//loadDict(whiteSuffixTrie, "add", "config/sensitive/white/suffix")
		loadDict(whiteSuffixTrie, "add", "config/sensitive")
	}
	return whiteSuffixTrie
}

// ClearWhiteSuffixTrie 清空白名单后缀Trie树
func ClearWhiteSuffixTrie() {
	whiteSuffixTrie = NewTrie()
}

func loadDict(trieHandle *Trie, op, path string) {

	var loadAllDictWalk filepath.WalkFunc = func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		initTrie(trieHandle, op, path)

		return nil
	}

	err := filepath.Walk(path, loadAllDictWalk)
	if err != nil {
		panic(err)
	}
}

func initTrie(trieHandle *Trie, op, path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("fail to open file %s %s", path, err.Error()))
	}

	defer f.Close()

	log.Infof("%s Load dict: %s", time.Now().Local().Format("2006-01-02 15:04:05 -0700"), path)

	buf := bufio.NewReader(f)
	for {
		line, isPrefix, e := buf.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}

		if word := strings.TrimSpace(string(line)); word != "" {
			tmp := strings.Split(word, " ")
			s := strings.Trim(tmp[0], " ")

			if "add" == op {
				trieHandle.Add(s)

			} else if "del" == op {
				trieHandle.Del(s)
			}
		}
	}

	return
}
