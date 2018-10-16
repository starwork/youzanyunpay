package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//Signature sha1签名
func Signature(params ...string) string {
	//sort.Strings(params)
	//h := sha1.New()
	//for _, s := range params {
	//	io.WriteString(h, s)
	//}
	//return fmt.Sprintf("%x", h.Sum(nil))

	h := md5.New()
	//h.Write([]byte("123456")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	return hex.EncodeToString(cipherStr)
}
