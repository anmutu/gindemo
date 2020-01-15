/*
  author='du'
  date='2020/1/15 20:09'
*/
package main

import "github.com/gin-gonic/gin"

//G:\lanuage_category\go\src\gindemo>curl http://127.0.0.1:8080/du
//{"message":"dududu"}

func main() {
	r := gin.Default()
	r.GET("/du", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "dududu",
		})
	})
	r.Run()
}
