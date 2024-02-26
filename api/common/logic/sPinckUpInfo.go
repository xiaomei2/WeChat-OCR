package logic

import "github.com/zeromicro/x/errors"

func SPickUpInfo(imagescode string) (resultMap map[string]string, err error) {
	resultMap = make(map[string]string)
	//第一个文本
	text, err := ExtractText(imagescode)
	if err != nil {
		return nil, err
	}
	var words = []string{"顺丰快递", "EMS快递", "申通快递", "圆通快递", "韵达快递", "极兔速递", "中通快递", "尼乔卫生站", "邮政快递", "韵达快递", "圆通速递"}
	var words2 = []string{"取件码", "取件码："}
	//第一个map
	//第二文本
	newText, one := GetNewTestAndCreateMap(text, words)
	res := IsKeywordBefore(text, words, words2)
	if res {
		//在前面
		two, err := GetKeyAndContentMapForFront(newText)
		if err != nil {
			return nil, errors.New(1001, "解析失败")
		}
		for key := range one {
			if value, ok := two[key]; ok {
				resultMap[value] = one[key]
			}
		}
	} else {
		two, err := GetKeyAndContentMapForAfter(newText)
		if err != nil {
			return nil, errors.New(1001, "解析失败")
		}
		for key := range one {
			if value, ok := two[key]; ok {
				resultMap[value] = one[key]
			}
		}
	}
	return resultMap, nil
}
