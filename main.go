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

func main() {
	watcher = services.NewFileWatcher()
	if err := watcher.Open("test.log"); err != nil {
		fmt.Println("Error opening watcher:", err)
		return
	}
	defer watcher.Close()
	logInfo, err := watcher.GetOneNewLine()
	if err != nil {
		fmt.Println("Error get log")
	}
	fmt.Println("Log infoxx:", logInfo)
	configer, _ := conf.ParseJsonConfig()
	fmt.Println("Configer:", configer)
	r := gin.Default()
	// add watcher to gin
	r.Use(func(c *gin.Context) {
		c.Set("watcher", watcher)
		c.Next()
	})

	// Example GET request
	r.GET("/ping", func(c *gin.Context) {
		w := c.MustGet("watcher").(*services.FileWatcher)
		logInfo, err := w.GetOneNewLine()
		if err != nil {
			fmt.Println("Error get log")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": logInfo,
		})
	})
	port, _ := configer.GetPort()
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
