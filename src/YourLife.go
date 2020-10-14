package main

import "fmt"

type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
}

func GuardarEntrada(tit string, d string, h string, text string){
	var n Entrada
	n.Titulo = tit
	n.Dia = d
	n.Hora = h
	n.Texto = text
}

func main() {

	var e1 Entrada

	menu := `
	Bienvenido a YourLife
	[ 1 ] Saludo
	[ 2 ] Introducir una entrada
	[ 3 ] Leer una entrada
    `
	fmt.Print(menu)

	var opcion int 
	fmt.Scanln(&opcion)

	switch opcion{
			case 1:
				fmt.Println("Bienvenido a YourLife")
			case 2:
				fmt.Println("Ponle titulo a tu dia")
				fmt.Scanln(&e1.Titulo)
				fmt.Println("¿Que dia es hoy?")
				fmt.Scanln(&e1.Dia)
				fmt.Println("¿Que hora es hoy?")
				fmt.Scanln(&e1.Hora)
				fmt.Println("Cuentanos que te ha ocurrido hoy")
				fmt.Scanln(&e1.Texto)
	}

	GuardarEntrada(e1.Titulo,e1.Dia,e1.Hora,e1.Texto)

}







