package main

import "fmt"

type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
}

func main() {
	var e Entrada
	e.Titulo = "\nPrimera entrada del diario\n"
	e.Dia = "13/10/2020\n"
	e.Hora = "19:00\n"
	e. Texto = "Escribo la primera entrada del diario.\n"

    fmt.Println(e)
}







