package util

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"os/signal"
	"syscall"
)

//WaitForExit 挂起
func WaitForExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-ch
}

// GenStrEncodedAndSalt 对指定字符串进行加密, 返回加密后的字符串和盐
func GenStrEncodedAndSalt(str *string) (string, string) {
	salt := "static salt"
	hash := sha1.New()
	hash.Write([]byte(*str + salt))
	credential := hex.EncodeToString(hash.Sum(nil))
	return credential, salt
}
