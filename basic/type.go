package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//内置基础内型
	//布尔类型：bool：true , false
	//整型：int8 byte int16 int unit unitptr (uintptr是整型,可以足够保存指针的值得范围)
	//浮点类型：float32 float64
	//复数类型：complex64 complex128
	//字符串：string
	//字符类型：rune
	//错误类型：error
	//复合类型：指针（pointer） 数组（array）切片（slice）字典（map）通道（chan）结构体（struct）接口（interface）

	//普通类型转换
	//string 转 int
	str := "30k"
	intValue, _ := strconv.Atoi(str)
	fmt.Println(reflect.TypeOf(intValue))

	//string 转 int64 或者int8等
	int64Value, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(reflect.TypeOf(int64Value))

	//int转string
	intTmp := 100
	strTmp := strconv.Itoa(intTmp)
	fmt.Println(reflect.TypeOf(strTmp))

	//int64转string
	var intTmp64 int64
	intTmp64 = 0xA
	strTmp = strconv.FormatInt(intTmp64, 10)
	fmt.Println(reflect.TypeOf(strTmp))

	jsonStr := "{\"euin\":\"342d05ad579b8e068fdc29f30384c9b3\",\"s\":\"o\",\"videolst\":[{\"ctime\":\"1970-01-01 08:00:00\",\"cull\":0,\"desc\":\"假如生活捉弄了你...不要悲伤...不要心急...\\r\\n反正...以后也不会好过...\",\"duration\":\"03:25\",\"pic\":\"http://vpic.video.qq.com/50350981/l0553wqx9ar_160_90_3.jpg\",\"play_count\":\"5.6万\",\"title\":\"今天不开心没关系，反正明天也不会好过\",\"title_s\":\"今天不开心没关系，反正明天也不会好过\",\"uploadtime\":\"2017-09-21\",\"url\":\"https://v.qq.com/x/page/l0553wqx9ar.html\",\"vid\":\"l0553wqx9ar\"},{\"ctime\":\"1970-01-01 08:00:00\",\"cull\":0,\"desc\":\"\",\"duration\":\"03:20\",\"pic\":\"http://vpic.video.qq.com/51661863/w14216higcw_160_90_3.jpg\",\"play_count\":\"7972\",\"title\":\"假如生活捉弄了你...不要悲伤...不要心急...反正...以后也不会好过...\",\"title_s\":\"假如生活捉弄了你...不要悲伤...不要心急...反正...以后也不会好过...\",\"uploadtime\":\"2017-09-21\",\"url\":\"https://v.qq.com/x/page/w14216higcw.html\",\"vid\":\"w14216higcw\"}],\"vtotal\":203}"

	//关于这样的字符串如何通过某个字段判断是否获取到数据或者其他，这里就是涉及到interface 解析 成map 或者string 涉及到断言。普通类型的直接强转就行，涉及到复杂的结构就不要使用断言。

	//为啥要这么定义，因为后面的不是string只能用interface{}来表示任何或者未知类型
	var jsonMap map[string]interface{} = make(map[string]interface{})

	//首先json_decode
	json.Unmarshal([]byte(jsonStr), &jsonMap)

	//获取videolst的里面的url写到slice里面
	var urls []string
	fmt.Println(jsonMap["videolst"])
	if a, ok := jsonMap["videolst"].([]interface{}); ok {
		for _, b := range a {
			if c, ok := b.(map[string]interface{}); ok {
				for k2, d := range c {
					if e, ok := d.(string); ok {
						if k2 == "url" {
							urls = append(urls, e)
						} else {
							continue
						}
					}
				}
			}
		}
	}
	//上面通过多次断言获取了url的列表
	fmt.Println(urls)

}