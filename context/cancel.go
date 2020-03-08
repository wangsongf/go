package main

import (
	"context"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	//下面是1秒后超时
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//下面是当前时间+1秒取消
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1))
	doOneWork(ctx)
	defer cancel()
	fmt.Println(ctx.Err())
}

func doOneWork(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop \n")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Println(incr)
		}
	}
}