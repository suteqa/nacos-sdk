package example

import (
	"fmt"
	"github.com/suteqa/nacos-sdk/clients/naming_client"
	"github.com/suteqa/nacos-sdk/utils"
	"github.com/suteqa/nacos-sdk/vo"
	"log"
)


func RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := client.RegisterInstance(param)
	if success {
		log.Fatalf("[INFO] 服务名 [%s] 注册成功  address [%s:%d] \n", param.ServiceName, param.Ip, param.Port)
	} else {
		log.Fatalf("[ERROR] 服务名 [%s] 注册失败  address [%s:%d] ERROR=%v\n", param.ServiceName, param.Ip, param.Port,err)
	}
}

func ExampleServiceClient_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Println(success)
}

func ExampleServiceClient_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Println(success)
}

func ExampleServiceClient_GetService(client naming_client.INamingClient) {
	service, _ := client.GetService(vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"a"},
	})
	fmt.Println(utils.ToJsonString(service))
}

func ExampleServiceClient_Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Subscribe(param)
}

func ExampleServiceClient_UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Unsubscribe(param)
}
