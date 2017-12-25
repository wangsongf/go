package main
import "github.com/garyburd/redigo/redis"
import "fmt"

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed, err:", err)
        return
    }
     _, err = c.Do("Set", "name", "nick")
    if err != nil {
        fmt.Println(err)
        return
    }
    //set & get
    r, err := redis.String(c.Do("Get", "name"))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(r)
    //mset & mget

    //批量设置
    _, err = c.Do("MSet", "name", "nick", "age", "18")
    if err != nil {
        fmt.Println("MSet error: ", err)
        return
    }

    r2, err := redis.Strings(c.Do("MGet", "name", "age"))
    if err != nil {
        fmt.Println("MGet error: ", err)
        return
    }
    fmt.Println(r2)

    //hset & hget  hash操作
    _, err = c.Do("HSet", "names", "nick", "suoning")
    if err != nil {
        fmt.Println("hset error: ", err)
        return
    }

    r, err = redis.String(c.Do("HGet", "names", "nick"))
    if err != nil {
        fmt.Println("hget error: ", err)
        return
    }
    fmt.Println(r)


    defer c.Close()
}