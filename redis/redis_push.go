package main
import "github.com/garyburd/redigo/redis"
import "fmt"

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed, err:", err)
        return
    }
    //设置过期时间
    _, err = c.Do("expire", "names", 5)
    if err != nil {
        fmt.Println("expire error: ", err)
        return
    }

    // 队列  lpush & lpop & llen
    _, err = c.Do("lpush", "Queue", "nick", "dawn", 9)
    if err != nil {
        fmt.Println("lpush error: ", err)
        return
    }
    for {
        r1, err := redis.String(c.Do("lpop", "Queue"))
        if err != nil {
            fmt.Println("lpop error: ", err)
            break
        }
        fmt.Println(r1)
    }
    r3, err := redis.Int(c.Do("llen", "Queue"))
    if err != nil {
        fmt.Println("llen error: ", err)
        return
    }
    fmt.Println(r3)


    defer c.Close()
}