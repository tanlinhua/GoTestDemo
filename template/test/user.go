package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfoStruct struct {
	Id   int
	Age  int
	Name string
	Addr string
}

func UserInfo(ctx *gin.Context) {
	user_info := UserInfoStruct{
		Id:   101,
		Age:  18,
		Name: "张三",
		Addr: "北京市"}
	ctx.HTML(http.StatusOK, "user/userinfo.html", user_info)
}

func ArrController(ctx *gin.Context) {
	// 数组渲染
	arr := [3]int{1, 3, 5}
	ctx.HTML(http.StatusOK, "user/arr.html", arr)

}

func ArrStruct(ctx *gin.Context) {
	// 结构体数据
	arr_struct := [3]UserInfoStruct{
		{Id: 102, Age: 18, Name: "张三", Addr: "北京"},
		{Id: 103, Age: 19, Name: "李四", Addr: "上海"},
		{Id: 104, Age: 20, Name: "王五", Addr: "广州"},
	}

	ctx.HTML(http.StatusOK, "user/arr_struct.html", arr_struct)
}

func MapController(ctx *gin.Context) {
	// map数据结构
	map_data := map[string]string{
		"name": "张三",
		"age":  "20"}

	ctx.HTML(http.StatusOK, "user/map.html", map_data)
}

func MapStruct(ctx *gin.Context) {
	// map数据结构
	map_struct := map[string]UserInfoStruct{
		"101": {Id: 101, Name: "张三", Addr: "北京"},
		"102": {Id: 102, Name: "李四", Addr: "北京", Age: 19},
		"103": {Id: 103, Name: "王五", Addr: "北京", Age: 20}}

	ctx.HTML(http.StatusOK, "user/map_struct.html", map_struct)
}
