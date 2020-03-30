package main

import (
	"gin-im/router"
)

func main() {
	webIm := router.InitRouter()
	_ = webIm.Run(":8008")
}
