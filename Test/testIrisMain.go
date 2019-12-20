package main

import (
	"github.com/kataras/iris"
)

func main()  {
	app:=iris.New()
	app.Get("/", func(context iris.Context) {
		context.HTML("<b>hello1111 Iris</b>")
	})
	//app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog:false}))
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
}
