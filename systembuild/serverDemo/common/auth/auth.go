package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"serverDemo/common/basic"
	"serverDemo/common/dstruct"
	"serverDemo/common/log"

	"github.com/gin-gonic/gin"
)

var secret = "0a113ef6b61820daa5611c870ed8d5ee"

func init() {
	if len(basic.App.Secret) > 10 {
		secret = basic.App.Secret
	}
}

func AesDecrypt(token string) string {
	tokenByte, err := hex.DecodeString(token)
	if err != nil {
		return ""
	}
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		//fmt.Println("err is:", err)
		return ""
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(tokenByte))
	blockMode.CryptBlocks(origData, []byte(tokenByte))
	origData = PKCS5UnPadding(origData)
	return string(origData)
}

func AesEncrypt(src string) string {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		//fmt.Println("key error1", err)
		log.Error(err)
		return ""
	}
	//src := fmt.Sprintf("%v|%v|%v", uid, time.Now().UnixNano()/1000000, areaId)
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return hex.EncodeToString(crypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func ParseHead(c *gin.Context) (*dstruct.ReqHead, error) {
	reqHead := &dstruct.ReqHead{}
	if err := c.BindQuery(reqHead); err != nil {
		log.Warn(err)
		return nil, err
	}
	return reqHead, nil
}

func CalSign() (string, error) {
	return "", nil
}
