package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "This is an main page...",
	})
}

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", HomepageHandler)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
