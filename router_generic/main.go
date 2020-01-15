/*
  author='du'
  date='2020/1/15 22:14'
*/
package main

import "github.com/gin-gonic/gin"

//所有以user开关的请求都会到"may i have ur name?"这里来。
//curl "http://127.0.0.1:8080/user/xxx
// may i have ur name?
func main() {
	r := gin.Default()
	r.GET("/user/*action", func(context *gin.Context) {
		context.String(200, " may i have ur name?")
	})
	r.Run()
}
