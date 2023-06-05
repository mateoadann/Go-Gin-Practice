package main

import "github.com/gin-gonic/gin"

func main() {
	//Crear un router con GIN
	router := gin.Default()

	//Captura la solicitud GET "/hello-world"
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO WORLD, MATEO HERE!!!",
		})
	})

	router.Run()

}