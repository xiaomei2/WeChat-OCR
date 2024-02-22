package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"wechat-ocr/common/logic"
)

type wuliu struct {
	content string
	key     string
}

func main() {
	text := getImagesText()
	var words = []string{"顺丰快递", "EMS快递", "申通快递", "圆通快递", "韵达快递", "极兔速递", "中通快递", "尼乔卫生站", "邮政快递", "韵达快递", "圆通速递"}
	newText, one := logic.GetNewTestAndCreateMap(text, words)
	two, err := logic.GetKeyAndContentMapForFront(newText)
	if err != nil {
		panic(err)
	}
	resultMap := map[string]string{}
	for key := range one {
		if value, ok := two[key]; ok {
			resultMap[value] = one[key]
		}
	}
	fmt.Printf("newKeywordMap:%v\n", one)
	fmt.Printf("前面结果：%v\n", two)
	fmt.Printf("最终结果:%v\n", resultMap)
	//result, err := logic.GetKeyAndContentMapForAfter(text)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("newText:%v\n", newText)
	//fmt.Printf("后面结果：%v\n", result)
	//fmt.Printf("中间结果：%v\n", resultMap)
}
func getImagesText() string {
	// 读取本地图片文件
	fileData, err := ioutil.ReadFile("./images/7.jpg")
	if err != nil {
		panic(err)
	}
	// 将图片内容进行 Base64 编码
	imageBase64 := base64.StdEncoding.EncodeToString(fileData)
	text, err := logic.ExtractText(imageBase64)
	if err != nil {
		panic(err)
	}
	return text
}
