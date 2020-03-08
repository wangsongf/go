package main

import   "fmt"
import   "io"
import   "log"
import   "net/http"
import   "github.com/julienschmidt/httprouter"

// httprouter.Params 是匹配到的路由参数，比如规则 /user/:id 中 的 :id 的对应值
func handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "hello, httprouter")
}

func main() {
    router := httprouter.New()
    router.GET("/", handle)

    if err := http.ListenAndServe(":12345", router); err != nil {
        fmt.Println("start http server fail:", err)
    }
}