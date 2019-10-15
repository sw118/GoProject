package main

import (
	"GoProject/common"
	"fmt"
)

func main() {
	//common.Request() //http://10.104.28.27:8989/adam
	resut := common.RunMain("GET", "http://www.myurl.com?cid=__CID__&imei=__IMEI__&os=__OS__&timestamp=__TS__&plan=__PLAN__&unit= __UNIT__&callback_url=__CALLBACK_URL__", nil)
	fmt.Println(resut.Body)

}
