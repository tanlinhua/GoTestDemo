package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "swagger.demo/docs" // 执行swag init生成的docs文件夹路径 _引包表示只执行init函数
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	disablingKey := "SWAGGER_DISABLE"

	debug := false
	if debug {
		os.Setenv(disablingKey, "true") // 禁用swagger
	}

	router := gin.Default()
	// http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, disablingKey))

	router.GET("/test", Test)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/print", Print)
		v1.GET("/hello", Hello)
	}
	router.Run()
}

// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router /api/v1/print [get]
func Print(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": name,
	})
}

// @Summary Hello接口
// @Description Hello接口
// @Tags 用户信息
// @Success 200 {string} json "{"message":"success"}"
// @Router /api/v1/hello [get]
func Hello(c *gin.Context) {
	// 当响应码为200时，返回JSON格式数据
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// @Summary 测试接口
// @Description 描述信息
// @Success 200 {string} string  "ok"
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(200, "ok")
}
