package springcloud

import (
	"github.com/suteqa/nacos-sdk/clients"
	"github.com/suteqa/nacos-sdk/clients/naming_client"
	"github.com/suteqa/nacos-sdk/common/constant"
	"github.com/suteqa/nacos-sdk/vo"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Service(configs []constant.ServerConfig, serverName string, serverPort uint64) {
	client, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": configs,
		"clientConfig": constant.ClientConfig{
			TimeoutMs:      9000,
			ListenInterval: 10000,
			//CacheDir:            "data/nacos/cache",
			NotLoadCacheAtStart: true,
		},
	})
	ip := getIpAddr()
	RegisterServiceInstance(client, vo.RegisterInstanceParam{
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

func RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	if success {
		log.Printf("[INFO] 服务名 [%s] 注册成功  address [%s:%d] \n", param.ServiceName, param.Ip, param.Port)
	} else {
		log.Fatalf("[ERROR] 服务名 [%s] 注册失败  address [%s:%d] \n", param.ServiceName, param.Ip, param.Port)
	}
	go func() {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-exitChan
		log.Printf("[EXIT] 服务关闭 [%s]  address [%s:%d] \n", param.ServiceName, param.Ip, param.Port)
		_, _ = client.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          param.Ip,
			Port:        param.Port,
			Cluster:     param.ClusterName,
			ServiceName: param.ServiceName,
			GroupName:   param.GroupName,
			Ephemeral:   true, //立刻删除服务
		})
		os.Exit(1)
	}()

}

func DeRegisterServiceInstance() {

}
