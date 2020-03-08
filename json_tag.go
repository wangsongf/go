package main

import (
	"fmt"
	"encoding/json"
)

type UserInfo struct {
	UsrId int64      `json:"user_id,omitempty"`
	NickName string  `json:"nickname"`
	Address string   `json:"-" `
}

func main () {
	var u UserInfo = UserInfo{
		NickName: "zhanglinpeng",
		Address: "",
	}
	rl, err:= json.Marshal(u)
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}
	// os.Stdout.Write(rl)
	fmt.Println(string(rl))
	var vlr UserInfo
	err1 := json.Unmarshal(rl, &vlr)
	if err != nil {
		fmt.Println("json unmarshal error: ", err1)
	}
	fmt.Printf("%v\n", vlr)
}