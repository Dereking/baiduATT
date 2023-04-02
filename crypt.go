package main

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算文本的md5值
func SumString(content string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(content))
	bys := md5Ctx.Sum(nil)
	//bys := md5.Sum([]byte(content))//这个md5.Sum返回的是数组,不是切片哦
	value := hex.EncodeToString(bys)
	return value
}
