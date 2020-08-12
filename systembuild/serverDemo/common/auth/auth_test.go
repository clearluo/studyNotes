package auth

import (
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	token := `0a0d13323b3c5ef3a28b5cd3b33ea3526c14879ef9aa0c2ba263e84ffa832843`
	str := AesDecrypt(token)
	t.Log(str)
	fmt.Println(str)
}
