package springcloud

import (
	"github.com/suteqa/nacos-sdk/clients"
	"github.com/suteqa/nacos-sdk/common/constant"
	"github.com/suteqa/nacos-sdk/example"
	"github.com/suteqa/nacos-sdk/vo"
	"net"
)

func InitService(addr, serverName string, port, serverPort uint64) {
	client, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: addr,
				Port:   port,
			},
		},
		"clientConfig": constant.ClientConfig{
			TimeoutMs:      9000,
			ListenInterval: 10000,
			//CacheDir:            "data/nacos/cache",
			NotLoadCacheAtStart: true,
		},
	})
	ip := getIpAddr()
	example.RegisterServiceInstance(client,vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        serverPort,
		ServiceName: serverName,
		Weight:      1,
		ClusterName: "DEFAULT",
		Enable:      true,
		Healthy:     true,
		Metadata: map[string]string{
			"preserved.register.source": "SPRING_CLOUD",
		},
		Ephemeral: true,
	})

	select {

	}
}

func getIpAddr() string {
	address, _ := net.InterfaceAddrs()
	for _, address := range address {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}



func DeRegisterServiceInstance() {

}
