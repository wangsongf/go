package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//组合二维等长map,生成等长json 例如：{"people1":{"name":"zengzhihai"},"people2":{"name":"liudehua"}}
	//创建两个2位map，一个二维map，二维map只接受为map的值
	var mapList1 = make(map[string]string)
	var mapList2 = make(map[string]string)
	var mapList12 = make(map[string]map[string]string)
	mapList1["name"] = "zengzhihai"
	mapList2["name"] = "liudehua"
	//这里是二维map接受两个基本map的值
	mapList12["people2"] = mapList2
	mapList12["people1"] = mapList1
	json12, _ := json.Marshal(mapList12)
	fmt.Println(string(json12))

	//组合3维或者多维的map的json。例如：{"result":1,"data":{"list":[{"titltid":"807242","title":"屌丝男士-饿了吃肉"},{"titltid":"8073342","title":"屌丝男士-波多来了"}],"page":1,"limit":30,"count":1,"total":1}}
	//这里是一个非常复杂的例子，当然我们主要是讲如何使用map
	var mapTlist1 = make(map[string]string)
	mapTlist1 = map[string]string{"titltid": "807242", "title": "屌丝男士-饿了吃肉"}
	var mapTlist2 = make(map[string]string)
	mapTlist2 = map[string]string{"titltid": "8073342", "title": "屌丝男士-波多来了"}
	var mapTlist3 = make([]interface{}, 0)
	var mapTlist4 = make(map[string]interface{})
	var mapTlistLast = make(map[string]interface{})
	mapTlist3 = append(mapTlist3, mapTlist1)
	mapTlist3 = append(mapTlist3, mapTlist2)
	mapTlist4["page"] = 1
	mapTlist4["limit"] = 30
	mapTlist4["count"] = 1
	mapTlist4["total"] = 1
	mapTlist4["list"] = mapTlist3
	mapTlistLast["result"] = 1
	mapTlistLast["data"] = mapTlist4
	fmt.Println(mapTlistLast)
	jsonLast, _ := json.Marshal(mapTlistLast)
	fmt.Println(string(jsonLast))

}