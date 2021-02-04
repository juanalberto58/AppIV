package main

import(
	"github.com/gin-gonic/gin"
	"github.com/juanalberto58/AppIV/src/m"
)

var new m.Entrada

func introducirEntrada(c *gin.Context){
	
	titulo := c.PostForm("titulo")
	dia := c.PostForm("dia")
	hora := c.PostForm("hora")
	entrada := c.PostForm("entrada")

	new.GuardarEntrada(titulo,dia,hora,entrada)

	c.JSON(200, gin.H{
		"Mensaje": "Entrada añadida con exito",
		"Titulo": titulo,
		"Dia": dia,
		"Hora": hora,
		"Entrada": entrada,
	})
}

func obtenerEntrada(c *gin.Context){
	
	c.JSON(200, gin.H{
		"Mensaje": "La entrada es: ",
		"Titulo": new.GetTitulo(),
		"Dia": new.GetDia(),
		"Hora": new.GetHora(),
		"Entrada": new.GetTexto(),
	})
}


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

	r.POST("/anadeentrada", introducirEntrada)
	r.GET("/obtenerEntrada", obtenerEntrada)

	r.Run()
}


