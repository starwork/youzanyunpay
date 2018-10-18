package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendRequest(httpUrl string, method string, paramMap map[string]string, files interface{}) []byte {
	fmt.Println("request url", httpUrl)
	fmt.Println("request method", method)
	fmt.Println("request paramMap", paramMap)

	var body io.Reader
	v := url.Values{}
	for key, val := range paramMap {
		v.Set(key, val)
	}
	//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	body = ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest(method, httpUrl, body)
	if err != nil {
		panic(err)
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//处理返回结果
	response, _ := client.Do(reqest)

	//返回的状态码
	//status := response.StatusCode
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	fmt.Println("request resp", string(content))
	return content
}
