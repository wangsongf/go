 //go发送http的get请求

package main
  
import (  
    "fmt"  
    "io/ioutil" //io模块  
    "net/http"  //http服务模块  
)

func main() {  
    //go 发送http的get请求  
    //response, err := http.Get("http://www.baidu.com")
    response, err := http.Get("http://o.baidu.com")

    //如果发送失败
    if err != nil {
        //handle error
        fmt.Println("发送get请求报错了。")
    }

    //这里是析构函数关闭body请求
    defer response.Body.Close()

    //读取响应的body
    body, _ := ioutil.ReadAll(response.Body)

    //打印响应的body
    fmt.Println(string(body))
}
