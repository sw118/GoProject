package common

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RequestPost(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return request
}

func RequestGet(method string, url string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return request
}

func RunMain(method string, url string, body io.Reader) *http.Response {
	client := &http.Client{}
	if method == "GET" {
		response, _ := client.Do(RequestGet(method, url))
		return response
	} else {
		response, _ := client.Do(RequestPost(method, url, body))
		return response
	}
}

func JsonParse(body map[string]interface{}) *bytes.Reader { //解析json
	jsonBody, _ := json.Marshal(body)   // 将map解析未[]byte类型
	reader := bytes.NewReader(jsonBody) // 将解析之后的数据转为*Reader类型
	//fmt.Println(reflect.TypeOf(reader))
	return reader
}
