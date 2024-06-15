package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func GetOrderNo() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf(
		"%s%d",
		time.Now().Format("20060102150405"),
		r.Intn(100000),
	)
}

// BuildMd5 md5生成
func BuildMd5(str string) string {
	hashedPassword := md5.Sum([]byte(str))
	pwd := fmt.Sprintf("%x", hashedPassword)
	return pwd
}
