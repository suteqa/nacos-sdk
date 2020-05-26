package springcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"sort"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
			}).DialContext,
		},
	}
}

// 发送Get请求
func httpGet(url string, header http.Header) (body []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	request.Header = header
	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// 发送Post请求
func httpPost(url string, body interface{}, header http.Header) (data []byte, err error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyStr))
	request.Header = header
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}

// 发送Post请求
func httpPostForm(url string, body interface{}, header http.Header) (data []byte, err error) {
	t := reflect.TypeOf(body)
	v := reflect.ValueOf(body)

	dmap := make(map[string]interface{})
	keyList := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		keyList = append(keyList, t.Field(i).Name)
		dmap[t.Field(i).Name] = v.Field(i).Interface()
	}
	//排序
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		s := fmt.Sprintf("%s=%s&", k, fmt.Sprintf("%v", dmap[k]))
		buffer.WriteString(s)
	}

	request, err := http.NewRequest("POST", url, buffer)
	request.Header = header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}
