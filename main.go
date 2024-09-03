package main

import (
	"trygo/rest-api/controllers"
	"trygo/rest-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main ()  {

  router := gin.Default()

  models.ConnDatabase()

    router.GET("/", func (c *gin.Context)  {
      c.JSON(200, gin.H{
      "message" : "hallo world",
      })
    })

  router.Use(cors.Default())

  router.GET("/api/quotes", controllers.GetData)

  router.POST("/api/quotes", controllers.PostData)

  router.GET("/api/quotes/:id", controllers.GetDataById)

  router.PUT("/api/quotes/:id", controllers.PutData)

  router.DELETE("/api/quotes", controllers.DeleteData)

  router.Run(":3000")

}
