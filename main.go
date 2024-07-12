/*
 * @Author: yujianming yujianming@macrosan.com
 * @Date: 2024-03-29 16:55:03
 * @LastEditors: yujianming yujianming@macrosan.com
 * @LastEditTime: 2024-03-29 16:55:12
 * @FilePath: \go_project\gin_project\example.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"net/http"

	"conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.ParseJsonConfig()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8666") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
