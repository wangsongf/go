package main

import (
	"time"
	"fmt"
)

func main() {
	//打印当前时间
	//time.Local即本地时区, 取决于运行的系统环境设置, 优先取”TZ”这个环境变量, 然后取/etc/localtime, 都取不到就用UTC兜底.
	//Now returns the current local time.
	n := time.Now()
	//格式化常规日志格式，不显示时区
	t1 := n.Format("2006-01-02 15:04:05")
	fmt.Println(t1)

	//格式化常规日志格式，显示时区
	t2 := n.String()
	fmt.Println(t2)

	//获取时间戳
	fmt.Println(n.Unix())
	//获取时间，精确到
	fmt.Println(n.UnixNano())

	currentTime := "2019-06-01 12:04:01"
	//解析时间到time类型,UTC时区
	u, _ := time.Parse("2006-01-02 15:04:05", currentTime)
	fmt.Println(u)
	//解析获取到当前时区：time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	lU, _ := time.ParseInLocation("2006-01-02 15:04:05", currentTime, time.Local)
	fmt.Println(lU)

	// 判断 n 比 u是否大
	fmt.Println(n.After(u))

	//判断n比u是否小
	fmt.Println(n.Before(u))

	//判断两个时间是否相等
	fmt.Println(n.Equal(u))

	//获取当前时间是哪一年
	fmt.Println(n.Year())
	//获取当前时间是这一年的哪一天
	fmt.Println(n.YearDay())
	//获取月份,天，小时，分钟，秒，纳秒
	fmt.Println(n.Month())
	fmt.Println(n.Day())
	fmt.Println(n.Hour())
	fmt.Println(n.Minute())
	fmt.Println(n.Second())
	fmt.Println(n.Nanosecond())

	//获取当前时间一年后，一个月后，一天后
	fmt.Println(n.AddDate(1, 0, 0).String())
	fmt.Println(n.AddDate(0, 1, 0).String())
	fmt.Println(n.AddDate(0, 0, 1).String())

	//获取昨天
	fmt.Println(n.AddDate(0, 0, -1).String())
	//获取前天
	fmt.Println(n.AddDate(0, 0, -2).String())

	unixT := 1389058332
	//时间戳格式化时间,不显示时区
	fmt.Println(time.Unix(int64(unixT), 0).Format("2006-01-02 15:04:05"))
	//时间戳格式化时间,显示时区
	fmt.Println(time.Unix(int64(unixT), 0).String())

	//定时器功能，定时循环拿到数据，次功能和下面的代码功能一模一样
	tCount1 := 0
	tCount2 := 0
	for range time.Tick(1 * time.Second) {
		if tCount1 == 2 {
			break
		}
		fmt.Println(1212121)
		tCount1++
	}
	for {
		if tCount2 == 2 {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println(2121212)
		tCount2++
	}

	//NewTimer 定时器功能
	tTimer1 := time.NewTimer(1 * time.Second)
	expire := <-tTimer1.C
	fmt.Println(expire)
	//NewTicker 定时器功能
	tTimer2 := time.NewTicker(1 * time.Second)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker running...")
				time.Sleep(3 * time.Second)
			case stop := <-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return
				}
			}
		}
	}(tTimer2)
	time.Sleep(3 * time.Second)
	ch <- true
	close(ch)

	//关于tk timer
	//Go语言的定时器实质是单向通道，time.Timer结构体类型中有一个time.Time类型的单向chan
	//NewTimer 和 NewTicker 的主要差异是 NewTimer能定时调整之间的间隔时间，NewTicker是循环的处理数据。
}