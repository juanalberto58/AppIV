package main

import(
	"log"
	"github.com/gin-gonic/gin"
	"github.com/juanalberto58/AppIV/src/m"
)

var new m.Entrada

type recuentrada struct {
	Titulo string `json:"Titulo" binding:"required"`
	Dia string `json:"Dia" binding:"required"`
	Hora string `json:"Hora" binding:"required"`
	Entrada string `json:"Entrada" binding:"required"`
}

func introducirEntrada(c *gin.Context){

	titulo := c.PostForm("titulo")
	dia := c.PostForm("dia")
	hora := c.PostForm("hora")
	entrada := c.PostForm("entrada")

	new.GuardarEntrada(titulo,dia,hora,entrada)

	uri := "/entrada/" + titulo
	c.Header("Location", uri)
	c.JSON(201, gin.H{
		"Mensaje": "Entrada añadida con exito",
		"Titulo": titulo,
		"Dia": dia,
		"Hora": hora,
		"Entrada": entrada,
	})
}


func introEntrada(c *gin.Context){

	var entr recuentrada
	err := c.ShouldBindJSON(&entr)

	if err != nil{
		c.JSON(400,gin.H{"Mensaje":"Error en la petición"})
		return
	}

	new.GuardarEntrada(entr.Titulo,entr.Dia,entr.Hora,entr.Entrada)

	uri := "/entrada/" + entr.Titulo
	c.Header("Location", uri)
	c.JSON(201, gin.H{
		"Mensaje": "Entrada añadida con exito",
		"Titulo": entr.Titulo,
		"Dia": entr.Dia,
		"Hora": entr.Hora,
		"Entrada": entr.Entrada,
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
		"status": "OK",
	})
}

func index(c *gin.Context){
	
	c.JSON(200,gin.H{
		"Mensaje": "El despliegue en heroku se ha realizado correctamente",
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
	r.POST("/entrada", introEntrada)

	r.GET("/entrada/:titulo", obtenerEntrada)
	r.GET("/obtenerEntrada", obtenerEntrada)

	r.POST("/modificarEntrada", editarEntrada)
	r.GET("/numeroEntradas", obtenerNumEntradas)
	r.GET("/status", stat)
	r.GET("/", index)

	return r
}



func main() {
	server().Run()
}


