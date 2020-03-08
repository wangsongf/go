
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if canceled(ctx) {
					break
				}
				doOneWork()
			}
			fmt.Println(i, "canceled.")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 5)
}

func canceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		fmt.Println("取消了")
		return true
	default:
		return false
	}
}

func doOneWork() {
	fmt.Println(111)
}