package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	logic2 "wechat-ocr/api/common/logic"
)

type wuliu struct {
	content string
	key     string
}

func main() {
	text := getImagesText()
	reslut, err := logic2.SPickUpInfo(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(reslut)
}
func getImagesText() string {
	// 读取本地图片文件
	fileData, err := ioutil.ReadFile("./images/7.jpg")
	if err != nil {
		panic(err)
	}
	// 将图片内容进行 Base64 编码
	imageBase64 := base64.StdEncoding.EncodeToString(fileData)

	return imageBase64
}
