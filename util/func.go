package util

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//WaitForExit 挂起
func WaitForExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-ch
}

// GenStrEncodedAndSalt 对指定字符串进行加密, 返回加密后的字符串和盐
func GenStrEncodedAndSalt(str *string) (string, string) {
	salt := GenRandomStr(32)
	hash := sha1.New()
	hash.Write([]byte(*str + salt))
	credential := hex.EncodeToString(hash.Sum(nil))
	return credential, salt
}

const alphaNumMap = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenRandomStr 生成指定长度的随机字符串
func GenRandomStr(length int) string {
	var bytes = make([]byte, 2*length)
	var outBytes = make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	mapLen := len(alphaNumMap)
	for i := 0; i < length; i++ {
		outBytes[i] = alphaNumMap[(int(bytes[2*i])*256+int(bytes[2*i+1]))%(mapLen)]
	}
	return string(outBytes)
}

// LocalDateTime 本地化时间处理
func LocalDateTime(date time.Time) time.Time {
	// 格式化时间，由于Dao模型中，对应时间字段为time.Time，保持和实体对齐，所以这里暂且只本地化时间处理
	// date.Local().Format(constant.DefaultTimestampFormat)
	return date.Local()
}
