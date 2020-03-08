
package main

import (
    "encoding/json"  
    "fmt"  
)

type Server struct {  
    ServerName string  
    ServerIp   string  
}

type Serverslice struct {  
    Servers []Server  
}

func main() {  
    //解析一个json  
    var s Serverslice  
    str := `{"servers":[{"serverName":"GZZ_VPN","serverIP":"127.0.0.3"},{"serverName":"SZ_VPN","serverIP":"127.0.0.4"}]}`  
    //这里是json_decode  
    json.Unmarshal([]byte(str), &s)  
    fmt.Println(s)  

    //添加一个json
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIp: "127.0.0.1"})
    s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIp: "127.0.0.2"})

    //这里相当于json_encode
    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }

    fmt.Println(string(b))
}

