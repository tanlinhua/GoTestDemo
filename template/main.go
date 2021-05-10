package main

import (
	"gin_template_demo/test"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// 注册模板
	engine.LoadHTMLGlob("template/**/*")
	// 注册静态文件
	engine.Static("./static", "static") // (relativePath, root string): 路径, 文件夹名称

	// 注册路由
	engine.GET("/userinfo", test.UserInfo)
	engine.GET("/arr", test.ArrController)
	engine.GET("/arr_struct", test.ArrStruct)
	engine.GET("/map", test.MapController)
	engine.GET("/map_struct", test.MapStruct)

	engine.Run(":9000")
}

// https://blog.csdn.net/qq_35709559/article/details/109147037
// https://blog.csdn.net/zhaogaolongsina/article/details/77126941
// https://blog.csdn.net/zhaogaolongsina/article/details/82997066
