package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Users struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Book string `json:"book"`
}

func StartServer() {
	log.Println("Server start up")

	var users = []Users{
		{Id: "1", Book: "Blue Train", Name: "John Coltrane", Age: 56},
		{Id: "2", Book: "Jeru", Name: "Gerry Mulligan", Age: 17},
		{Id: "3", Book: "Sarah Vaughan and Clifford Brown", Name: "Sarah Vaughan", Age: 39},
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Список покупок книг",
		})
		c.HTML(http.StatusOK, "users.html", gin.H{
			"Users": users,
		})
	})
	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
