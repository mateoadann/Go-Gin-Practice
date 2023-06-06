package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Estructura para el producto
type Product struct {
	ID             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Precio         int    `json:"precio"`
	Stock          int    `json:"stock"`
	Codigo         int    `json:"codigo"`
	Publicado      bool   `json:"publicado"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}


func main() {
	
	//Crear un router con GIN
	router := gin.Default()
	


	//Captura la solicitud GET "/hello-world"
	router.GET("/productos", func(c *gin.Context) {
		// Leer el archivo products.json
		file, err := ioutil.ReadFile("products.json")
		if err != nil {
			c.JSON(500, gin.H{"error": "No se pudo leer el archivo"})
			return
		}

		// Decodificar el archivo JSON en una lista de productos
		var productos []Product
		err = json.Unmarshal(file, &productos)
		if err != nil {
			fmt.Println("acá esta el error")
			c.JSON(500, gin.H{"error": "No se pudo analizar el archivo JSON"})
			return
		}
		// Devolver la lista de productos en formato JSON
		c.JSON(200, productos)
	})

	router.Run(":8080")

}

/**func HandlerGin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HELLO WORLD, MATEO HERE!!!",
	})
}**/



