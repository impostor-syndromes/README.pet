package main

import (
	"net/http" // httpパッケージをインポート

	"README.pet/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 静的ファイルの配信
	r.Static("/assets", "./assets")

	r.GET("/show-svg", func(c *gin.Context) {
		// Content-Typeを設定
		c.Header("Content-Type", "image/svg+xml")

		// SVGを返す
		svg := pkg.GenerateSVG(6)
		c.String(http.StatusOK, svg)
	})
	r.Run()
}
