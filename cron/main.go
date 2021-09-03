package main

import (
	"demo/cron/repo"
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppkg/glog"
	"github.com/ppkg/kit"
)

func main() {

	defer glog.Flush()
	flag.Parse()

	e := echo.New()

	// web服务
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	engine := repo.GetEngine()
	// engine.Sync2(new(models.Bug))

	result, err := engine.Query("select * from bug limit 5")
	if err != nil {
		glog.Infof("查询sql出错： %s", err.Error())
	}

	// json, _ := json.Marshal(result)
	// glog.Infof("查询sql结果： %s", string(json))

	glog.Infof("查询sql结果： %s", kit.JsonEncode(result))

	// 定时任务
	// err := jobs.StartJob()
	// if err != nil {
	// 	return
	// }

	e.Logger.Fatal(e.Start(":8888"))

}
