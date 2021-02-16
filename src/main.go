package main

import(
	"log"
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


func editarEntrada(c *gin.Context){
	
	titulo := c.PostForm("titulo")
	dia := c.PostForm("dia")
	hora := c.PostForm("hora")
	entrada := c.PostForm("entrada")

	new.EditarEntrada(titulo,dia,hora,entrada)

	c.JSON(200, gin.H{
		"Mensaje": "Entrada editada con exito",
		"Titulo": titulo,
		"Dia": dia,
		"Hora": hora,
		"Entrada": entrada,
	})
}

func obtenerNumEntradas(c *gin.Context){
	
	c.JSON(200, gin.H{
		"Mensaje": "El número de entradas es: ",
		"Num Entradas": new.GetCont(),
	})
}


func stat(c *gin.Context){
	
	c.JSON(200,gin.H{
		status: "OK",
	})
}


func LogMid() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Next()
		estado := c.Writer.Status()
		peticion := c.Request.Method
		ip := c.ClientIP()

		log.Printf("Código de estado: %v,Tipo de petición: %s,IP del cliente: %s", estado, peticion, ip)
	}
}


func server() *gin.Engine {

	r := gin.Default()

	r.Use(LogMid())

	r.POST("/anadeEntrada", introducirEntrada)
	r.GET("/obtenerEntrada", obtenerEntrada)
	r.POST("/modificarEntrada", editarEntrada)
	r.GET("/numeroEntradas", obtenerNumEntradas)
	r.GET("/status", stat)
	return r
}



func main() {
	server().Run()
}


