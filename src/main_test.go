package main

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestIntroducirEntrada(t *testing.T){
	
	r := server()

	w := httptest.NewRecorder()

	param := strings.NewReader("titulo=pruebatitulo&dia=pruebadia&hora=pruebahora&entrada=pruebaentrada")
	resp, err := http.NewRequest("POST", "/anadeEntrada", param)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo %v", err)
	}

	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")


	r.ServeHTTP(w, resp)

	if w.Code != 200 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 200, w.Code)
	}

	q := "{\"Dia\":\"pruebadia\",\"Entrada\":\"pruebaentrada\",\"Hora\":\"pruebahora\",\"Mensaje\":\"Entrada añadida con exito\",\"Titulo\":\"pruebatitulo\"}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}
}

func TestObtenerEntrada(t *testing.T){
	r := server()

	w := httptest.NewRecorder()

	resp, err := http.NewRequest("GET", "/obtenerEntrada", nil)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo %v", err)
	}

	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")


	r.ServeHTTP(w, resp)

	if w.Code != 200 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 200, w.Code)
	}

	q := "{\"Dia\":\"pruebadia\",\"Entrada\":\"pruebaentrada\",\"Hora\":\"pruebahora\",\"Mensaje\":\"Entrada añadida con exito\",\"Titulo\":\"pruebatitulo\"}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}

}


/*
func TestEditarEntrada(c *gin.Context){
	
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

func TestObtenerNumEntradas(c *gin.Context){
	
	c.JSON(200, gin.H{
		"Mensaje": "El número de entradas es: ",
		"Titulo": new.GetCont(),
	})
}


func server() *gin.Engine {

	r := gin.Default()

	r.POST("/anadeEntrada", introducirEntrada)
	r.GET("/obtenerEntrada", obtenerEntrada)
	r.POST("/modificarEntrada", editarEntrada)
	r.GET("/numeroEntradas", obtenerNumEntradas)

	return r
}



func main() {
	server().Run()
}


*/

