package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		articles := getAllArticles()

		render(ctx, gin.H{
			"title":   "Home Page",
			"payload": articles,
		}, "index.html")
	}
}

func getArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		articleID := ctx.Param("article_id")
		id, err := strconv.Atoi(articleID)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		}

		article, err := getArticleByID(id)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		}

		render(ctx, gin.H{
			"title":   article.Title,
			"payload": article,
		}, "article.html")
	}
}

func render(ctx *gin.Context, data gin.H, templateName string) {
	fmt.Println(ctx.Request.Header.Get("Accept"))

	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		ctx.XML(http.StatusOK, data["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, data)
	}
}
