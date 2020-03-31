package main

import (
	"github.com/suteqa/nacos-sdk/clients"
	"github.com/suteqa/nacos-sdk/common/constant"
	"github.com/suteqa/nacos-sdk/example"
	"github.com/suteqa/nacos-sdk/vo"
)

func main() {
	client, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "data.vjes.com",
				Port:   8848,
			},
		},
		"clientConfig": constant.ClientConfig{
			TimeoutMs:           9000,
			ListenInterval:      10000,
			CacheDir:            "data/nacos/cache",
			NotLoadCacheAtStart: true,
		},
	})

	example.ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        9000,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "a",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

	select {}

}
