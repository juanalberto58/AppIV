package main

import(
	"github.com/gin-gonic/gin"
)

/*************************************************************/


type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
	cont int
}

func (e *Entrada)Contructor(){
	e.Titulo = ""
	e.Dia = ""
	e.Hora = ""
	e.Texto = ""
	e.cont = 0
}

func (e Entrada)getTitulo() string{
	return e.Titulo
}


func (e *Entrada)setTitulo(tit string){
	e.Titulo=tit
}

/*************************************************************/

var new Entrada

func IntroducirTitulo(c *gin.Context){
	//c.Request.ParseForm()

	titulo := c.PostForm("titulo")

	new.setTitulo(titulo)
	
	c.JSON(200, gin.H{
		"Mensaje": "Titulo añadido con exito",
		"Titulo": titulo,
	})
}

func ObtenerTitulo(c *gin.Context){
	
	c.JSON(200, gin.H{
		"Mensaje": "El titulo es: ",
		"Titulo": new.getTitulo(),
	})
}


func main() {
	r := gin.Default()

	r.POST("/anadetitulo", IntroducirTitulo)
	r.GET("/obtenertitulo", ObtenerTitulo)

	r.Run()
}


