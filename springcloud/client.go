package springcloud

import (
	"errors"
	"fmt"
	"github.com/suteqa/nacos-sdk/vo"
	"net/http"
)

//获取请求地址
func getUrl(serverName, path string) (string, error) {
	instance, err := app.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serverName,
	})
	if err != nil {
		return "", errors.New("服务获取失败！")
	}
	return fmt.Sprintf("http://%s:%d/%s", instance.Ip, instance.Port, path), nil
}

func Get(serverName, path string, header http.Header) (body []byte, err error) {
	path, err = getUrl(serverName, path)
	if err != nil {
		return
	}
	return httpGet(path, header)
}

func Post(serverName, path string, parameter interface{}, header http.Header) (body []byte, err error) {
	path, err = getUrl(serverName, path)
	if err != nil {
		return
	}
	return httpPost(path, parameter, header)
}

func PostForm(serverName, path string, parameter interface{}, header http.Header) (body []byte, err error) {
	path, err = getUrl(serverName, path)
	if err != nil {
		return
	}
	return httpPostForm(path, parameter, header)
}
