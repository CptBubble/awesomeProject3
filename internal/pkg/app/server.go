package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Book string `json:"book"`
}

func (a *Application) StartServer() {
	log.Println("Server start up")

	var users = []Users{
		{Book: "Blue Train", Name: "John Coltrane", Age: 56},
		{Book: "Jeru", Name: "Gerry Mulligan", Age: 17},
		{Book: "Sarah Vaughan and Clifford Brown", Name: "Sarah Vaughan", Age: 39},
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string
		if id != "" {
			log.Printf("id received %s\n", id)
			intID, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("can't convert id %v", err)
				c.Error(err)
				return
			}

			users, err := a.repo.GetPromoByID(uint(intID))
			if err != nil {
				log.Printf("can't get promo by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"Book": users.Book,
				"Name": users.Name,
				"Age":  users.Age,
			})
			return
		}
		create := c.Query("create")
		if create != "" {
			log.Printf("create received %s\n", create)
			createBool, err := strconv.ParseBool(create) // пытаемся привести это к чиселке
			if err != nil {                              // если не получилось
				log.Printf("can't convert create %v", err)
				c.Error(err)
				return
			}

			if createBool {
				a.repo.NewRandRecords()
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "create not true",
			})

			return
		}
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
