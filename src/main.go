package main

import(
	"github.com/gin-gonic/gin"
	"github.com/juanalberto58/AppIV/src/m"
)

var new m.Entrada


func introducirTitulo(c *gin.Context){

	titulo := c.PostForm("titulo")

	new.SetTitulo(titulo)
	
	c.JSON(200, gin.H{
		"Mensaje": "Titulo añadido con exito",
		"Titulo": titulo,
	})
}

func obtenerTitulo(c *gin.Context){
	
	c.JSON(200, gin.H{
		"Mensaje": "El titulo es: ",
		"Titulo": new.GetTitulo(),
	})
}


func main() {
	r := gin.Default()

	r.POST("/anadetitulo", introducirTitulo)
	r.GET("/obtenertitulo", obtenerTitulo)

	r.Run()
}


