package main

import (
	"AppApplication"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

type param struct {
	token string            `json:"token"`
	msg   map[string]string `json:"msg"`
}

func main() {
	comm := tars.NewCommunicator()
	applocationObj := new(AppApplication.Server)
	//comm.StringToProxy("pkg.appApplication.tarsObj@tcp -h 172.30.6.55 -p 10230 -t 60000", applocationObj) // 生产
	comm.StringToProxy("pkg.appApplication.tarsObj@tcp -h 192.168.91.52 -p 61569 -t 60000", applocationObj)  //开发
	//comm.StringToProxy("pkg.appApplication.tarsObj@tcp -h 172.31.6.59 -p 10230 -t 60000", applocationObj) //预发
	//comm.StringToProxy("pkg.appApplication.tarsObj@tcp -h 172.31.6.58 -p 10230 -t 60000", applocationObj)  //预发
	//comm.StringToProxy("pkg.application.tarsObj@tcp -h 127.0.0.1 -p 57036 -t 60000", applocationObj)
	//comm.StringToProxy("pkg.appApplication.tarsObj@tcp -h 192.168.92.52 -t 60000 -p 10230 ", applocationObj)  //测试
	paramsStr := `{
	   "data":{
			"appId":1804,
	"memberId":"151528",
	"type":"5",
	"page":"1",
	"pageSize":1000,
	"orderBy":{"favoriteId":"ASC"}
	   },
	   "header":{
	       "traceId":"1122498647565529088",
	       "version":"1.0.0"
	   }
	}`

	respStr := "s"
	resCode, err := applocationObj.Get("favorite.list", paramsStr, &respStr)
	//resCode, err := applocationObj.Put("app.version", paramsStr, &respStr)
	fmt.Println(err)
	fmt.Println(resCode, "AAAAAAA", respStr)


}

/** favorite.list  params */
/*
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"5",
	"page":"1",
	"pageSize":1000,
	"orderBy":{"favoriteId":"ASC"}
},
*/

/** favorite.cancel  params */
/*
例一：
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"5",
	"routeId":"3746,3745,3744"
},

例二：
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"1",
	"routeId":"3746,3745,3744"
},

*/

/** favorite.info  params */
/**
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"1",
	"favoriteId":"3744"
},
*/

/** favorite.edit params */
/**
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"5",
	"favoriteId":"3750",
	"favoriteTitle":"北京,上海,南京",
	"favoriteDesc": {
		"routeTitle":"北京,上海,南京",
		"routeTime":"24小时35分",
		"routeDistance":"1000km",
		"charges":100000,
		"stationNum":100,
		"thumbnail":""
	}
},
 */

/** favorite.add params */
/**
"data":{
	"appId":1804,
	"memberId":"151528",
	"type":"5",
	"favoriteTitle":"北京2,上海,南京",
	"favoriteDesc": {
		"routeTitle":"北京,上海,南京",
		"routeTime":"24小时35分",
		"routeDistance":"1000km",
		"charges":100000,
		"stationNum":100,
		"thumbnail":""
	}
},
 */


