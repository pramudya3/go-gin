package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.html", gin.H{"titel": "Home Page"})
	// })

	router.GET("/", showIndexPage())
	router.GET("/article/view/:article_id", getArticle())

	router.Run()
}
