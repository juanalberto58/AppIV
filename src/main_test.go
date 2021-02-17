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

	if w.Code != 201 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 201, w.Code)
	}

	q := "{\"Dia\":\"pruebadia\",\"Entrada\":\"pruebaentrada\",\"Hora\":\"pruebahora\",\"Mensaje\":\"Entrada añadida con exito\",\"Titulo\":\"pruebatitulo\"}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}
}

func TestIntroEntrada(t *testing.T){
	
	r := server()

	w := httptest.NewRecorder()

	param := strings.NewReader("{\"Titulo\":\"pruebatitulo\",\"Dia\":\"pruebadia\",\"Hora\":\"pruebahora\",\"Entrada\":\"pruebaentrada\"}")
	resp, err := http.NewRequest("POST", "/entrada", param)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo %v", err)
	}

	resp.Header.Add("Content-Type", "application/json")


	r.ServeHTTP(w, resp)

	if w.Code != 201 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 201, w.Code)
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

	q := "{\"Dia\":\"pruebadia\",\"Entrada\":\"pruebaentrada\",\"Hora\":\"pruebahora\",\"Mensaje\":\"La entrada es: \",\"Titulo\":\"pruebatitulo\"}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}

}



func TestEditarEntrada(t *testing.T){
	
	r := server()

	w := httptest.NewRecorder()

	param := strings.NewReader("titulo=pruebatituloedit&dia=pruebadiaedit&hora=pruebahoraedit&entrada=pruebaentradaedit")
	resp, err := http.NewRequest("POST", "/modificarEntrada", param)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo %v", err)
	}

	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")


	r.ServeHTTP(w, resp)

	if w.Code != 200 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 200, w.Code)
	}

	q := "{\"Dia\":\"pruebadiaedit\",\"Entrada\":\"pruebaentradaedit\",\"Hora\":\"pruebahoraedit\",\"Mensaje\":\"Entrada editada con exito\",\"Titulo\":\"pruebatituloedit\"}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}
}



func TestObtenerNumEntradas(t *testing.T){

	r := server()

	w := httptest.NewRecorder()

	resp, err := http.NewRequest("GET", "/numeroEntradas", nil)

	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo %v", err)
	}

	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")


	r.ServeHTTP(w, resp)

	if w.Code != 200 {
		t.Errorf("Se esperaba %v, se obtuvo %v", 200, w.Code)
	}

	q := "{\"Mensaje\":\"El número de entradas es: \",\"Num Entradas\":2}"

	if w.Body.String() != q {
		t.Errorf("Entrada incorrecta, se esperaba: %v , Se ha obtenido %v", q, w.Body.String())
	}
}



