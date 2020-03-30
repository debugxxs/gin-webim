package router

import (
	"gin-im/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func InitRouter() *gin.Engine {
	go app.Hub.Run()
	router := gin.Default()

	runPath := GetGoRunPath()
	router.LoadHTMLGlob(filepath.Join(runPath, "views/*.html"))
	router.Static("/statics", filepath.Join(runPath, "statics"))
	router.StaticFile("/favicon.ico", filepath.Join(runPath, "statics/favicon.ico"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", nil)
	})
	router.Any("/im", app.WsServer)
	return router
}

func GetGoRunPath() string {
	appPath, _ := os.Getwd()
	return appPath
}
