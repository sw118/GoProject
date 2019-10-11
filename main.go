package main

import (
	"GoProject/common"
	"fmt"
)

func main() {
	//common.Request() //http://10.104.28.27:8989/adam
	resut := common.RunMain("GET", "https://www.cnblogs.com/nyist-xsk/p/10550812.html", nil)
	fmt.Println(resut.StatusCode)
	va := make(map[string]interface{})
	va["name"] = "sunwei"
	va["age"] = 20
	fmt.Println(common.JsonParse(va))

	res := common.RunMain("POST", "https://www.cnblogs.com/nyist-xsk/p/10550812.html", common.JsonParse(va))
	fmt.Println(res.Body)

}
