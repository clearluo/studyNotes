package sensitive

import (
	"strings"

	"serverDemo/common/log"
	"serverDemo/common/sensitive/trie"
)

// 查找敏感词
func QueryWords(str string) (bool, []string, string) {
	//log.Debug("需要检测:", str)
	if str != "" {
		ok, keyword, newText := trie.GblackTrie.Query(str)
		if ok {
			log.Debug("##敏感词:", keyword)
			return ok, keyword, newText
		}
	}
	return false, []string{}, str
}

// 添加敏感词
func addBlackWords(words ...string) {
	for _, s := range words {
		trie.GblackTrie.Add(strings.Trim(s, " "))
	}
}

// 删除敏感词
func deleteBlackWords(words ...string) {
	for _, s := range words {
		trie.GblackTrie.Del(strings.Trim(s, " "))
	}

}

// 显示所有敏感词
func showBlackWords() []string {
	words := trie.GblackTrie.ReadAll()
	return words
}
