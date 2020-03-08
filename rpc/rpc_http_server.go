//http的rpc服务
package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
//go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码

type Params struct {
	Width, Height int
}

type Rect struct {
}

//求面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

//求长度
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	//注册一个rect服务
	rpc.Register(rect)
	//把服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}