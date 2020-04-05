package springcloud

import (
	"github.com/suteqa/nacos-sdk/clients"
	"github.com/suteqa/nacos-sdk/common/constant"
	"github.com/suteqa/nacos-sdk/vo"
)

func Config(configs []constant.ServerConfig, dataId string) (string, error) {
	// 可以没有，采用默认值
	clientConfig := constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 30 * 1000,
		BeatInterval:   5 * 1000,
	}
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": configs,
		"clientConfig":  clientConfig,
	})
	return configClient.GetConfig(vo.ConfigParam{
		DataId:   dataId,
		Group:    "DEFAULT_GROUP",
		OnChange: nil,
	})
}
