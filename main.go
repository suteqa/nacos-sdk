package main

import (
	"fmt"
	"github.com/suteqa/nacos-sdk/common/constant"
	"github.com/suteqa/nacos-sdk/springcloud"
	"net"
	"net/http"
	"time"
)

func main() {
	//nacos服务
	configs := []constant.ServerConfig{
		{
			IpAddr: "data.vjes.com",
			Port:   8848,
		},
	}

	//http端口
	var port uint64 = 8522

	//启动注册服务
	springcloud.Service(configs, "mytest", port)

	go func() {
		timer := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-timer.C:
				header:=make(http.Header)
				header.Set("Authentication","")
				bb, _ := springcloud.Get("Server-Mall", "/wx/merchant/checkEnter",header)
				fmt.Println(string(bb))
			}
		}
	}()

	//获取配置中心
	str, _ := springcloud.Config(configs, "server-mall-dev.yaml")
	//strings.NewReader(str)
	fmt.Println(str)

	//启动http服务
	l, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr)
		_, _ = writer.Write([]byte("hello"))
	})
	//监听服务
	_ = http.Serve(l, nil)
}
