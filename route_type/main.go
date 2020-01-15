/*
  author='du'
  date='2020/1/15 20:38'
*/
package main

import "github.com/gin-gonic/gin"

//其中delete的curl是这样：
//curl -X DELETE  http://127.0.0.1:8080/delete
//du=>delete

//https://github.com/anmutu/gin

func main() {

	r := gin.Default()
	r.GET("get", func(context *gin.Context) {
		context.String(200, "du=>get")
	})
	r.GET("post", func(context *gin.Context) {
		context.String(200, "du=>post")
	})
	r.Handle("DELETE", "/delete", func(context *gin.Context) {
		context.String(200, "du=>delete")
	})
	//any的请求方法,里面有8种请求方法
	r.Any("any", func(context *gin.Context) {
		context.String(200, "du=>any")
	})
	r.Run()
}
