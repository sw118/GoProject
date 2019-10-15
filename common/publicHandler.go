package common

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func RequestPost(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	check(err)
	return request
}

func RequestGet(method string, url string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	check(err)
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
	return reader
}

func DataBase(driver string, url string, sqlQuery string) {
	db, err := sql.Open(driver, url)
	check(err)
	defer db.Close()
	rows, err := db.Query(sqlQuery)
	check(err)
	for rows.Next() {
		var info Info
		err = rows.Scan(&info.Id, &info.Name, &info.Age)
		check(err)
		fmt.Println(info)
	}
	rows.Close()
}

func RedisHandler(driver string, address string) {
	conn, err := redis.Dial(driver, address)
	check(err)
	defer conn.Close()
	_, err = conn.Do("SET", "key", "adam")
	check(err)
}
