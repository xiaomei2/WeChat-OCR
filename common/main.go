package main

import (
	"fmt"
	"wechat-ocr/common/logic"
)

func main() {
	text := "AAAAABAAABAAAAAAAAAABAABAAA"
	words := []string{"B"}
	newText, newKeywordMap := logic.GetNewTestAndCreateMap(text, words)
	fmt.Printf("newText:%v\n", newText)
	fmt.Printf("newKeywordMap:%v\n", newKeywordMap)
	result, err := logic.GetKeyAndContentMapForAfter(text)
	if err != nil {
		fmt.Println(err)
	}
	re, err := logic.GetKeyAndContentMapForFront(newText)
	if err != nil {
		panic(err)
	}
	item := make(map[string]string)
	for k, v := range newKeywordMap {
		if reValue, ok := re[v]; ok {
			item[k] = reValue
		}
	}
	fmt.Printf("后面结果：%v\n", result)
	fmt.Printf("前面结果：%v\n", re)
	fmt.Printf("中间结果：%v\n", item)
}
