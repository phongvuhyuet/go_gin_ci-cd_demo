package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
var router *gin.Engine
func render(c *gin.Context, data gin.H, templateName string) {

  switch c.Request.Header.Get("Accept") {
  case "application/json":
    // Respond with JSON
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    // Respond with XML
    c.XML(http.StatusOK, data["payload"])
  default:
    // Respond with HTML
    c.HTML(http.StatusOK, templateName, data)
  }

}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	router.Run()
}