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
	"fmt"
	"net/http"

	"conf"
	"services"

	"github.com/gin-gonic/gin"
)

var watcher *services.FileWatcher
var reader *services.Reader

func main() {
	configer, _ := conf.ParseJsonConfig()
	fmt.Println("Configer:", configer)
	r := gin.Default()

	// ReadFile
	reader, err := services.NewReader("logs/engine.log")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer reader.Close()

	r.Use(func(c *gin.Context) {
		c.Set("filter", reader)
		c.Next()
	})

	r.GET("/filter/:pattern", func(c *gin.Context) {
		// 获取参数
		pattern := c.Param("pattern")
		fmt.Printf("/filter/ get param pattern:%+v\n", pattern)
		w := c.MustGet("filter").(*services.Reader)
		filterString, err := w.FilterLines(pattern)
		if err != nil {
			fmt.Println("Error filter lines:", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": filterString,
		})
	})
	// Set Port and run service.
	port, _ := configer.GetPort()
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
