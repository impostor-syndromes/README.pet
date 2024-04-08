package api

import (
	"net/http" // httpパッケージをインポート

	"README.pet/api/pkg"
	"github.com/gin-gonic/gin"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("apiMain", apiMain)
}

func apiMain(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 静的ファイルの配信
	router.Static("/assets", "./assets")

	router.GET("/show-svg", func(c *gin.Context) {
		username := c.Query("username")

		contributions, err := pkg.FetchContributions(username)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Content-Typeを設定
		c.Header("Content-Type", "image/svg+xml")

		// SVGを返す
		svg := pkg.GenerateSVG(contributions[0])
		c.String(http.StatusOK, svg)
	})
	router.ServeHTTP(w, r)
}