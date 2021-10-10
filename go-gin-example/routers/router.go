package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wzhanjun/go-gin-example/middlewares/jwt"
	"github.com/wzhanjun/go-gin-example/pkg/setting"
	"github.com/wzhanjun/go-gin-example/routers/api"
	v1 "github.com/wzhanjun/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 获取认证
	r.GET("/auth", api.GetAuth)

	// v1

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签
		apiv1.GET("/tags", v1.GetTags)
		// 添加标签
		apiv1.POST("/tags", v1.AddTag)
		// 修改标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
