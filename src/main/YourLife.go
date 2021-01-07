package main

import "fmt"


func main() {

	var e Entrada
	var input string
	/*ar titu string
	var hor string
	var tex string*/
	//var con int

	menu := `
	Bienvenido a YourLife
	[ 0 ] Pulse 0 para salir
	[ 1 ] Introducir Titulo
	[ 2 ] Introducir el Dia
	[ 3 ] Introducir la Hora
	[ 4 ] Introducir la Entrada
	[ 5 ] Ver Entrada Completa
	[ 6 ] Calcular Número de Entradas Insertadas
    `
	//fmt.Print(menu)
	var cont int
	//var contenido [5]Entrada
	var opcion int 
	var exit bool
	exit=false
	

	//fmt.Scanln(&opcion)
	cont=0
	for{
		fmt.Print(menu)
		fmt.Scanln(&opcion)
		switch opcion{
			case 0:
				 exit=true
			case 1:
				exit=false
				fmt.Println("Ponle titulo a tu dia")
				fmt.Scanln(&input)
				e.setTitulo(input)
			case 2:
				exit=false
				fmt.Println("¿Que dia es hoy?")
				fmt.Scanln(&input)
				e.setDia(input)
			case 3:
				exit=false
				fmt.Println("¿Que hora es hoy?")
				fmt.Scanln(&input)
				e.setHora(input)
			case 4:
				exit=false
				fmt.Println("Cuentanos que te ha ocurrido hoy")
				fmt.Scanln(&input)
				e.setTexto(input)
				cont++
				e.setCont(cont)
			case 5:
				exit=false
				fmt.Println("Parece que has tenido un dia apasionante...")
				fmt.Println("Tu dia se titula... " + e.getTitulo() +"")
				fmt.Println("El dia para recordar es el... " + e.getDia() + "")
				fmt.Println("Lo escribistes a las... " + e.getHora() + "")
				fmt.Println("Esto fué lo que te paso... " + e.getTexto() + "")
			/*case 6:
				exit=false
				fmt.Println("¿De que dia quieres buscar la entrada?")
				fmt.Scanln(&input)
				titu, hor, tex = searchEntrada(contenido,input,cont)
				fmt.Println("Tu dia se titula... " + titu +"")
				fmt.Println("Lo escribistes a las... " + hor + "")
				fmt.Println("Esto fué lo que te paso... " + tex + "")*/
			case 6:
				exit=false
				fmt.Println("El número de entradas insertadas es  " , e.getCont() , "")
		}

		//contenido[cont] = e
		
		if(exit){
			break
		}
	}
	fmt.Println("Adiossss")

	//GuardarEntrada(e.Titulo,e.Dia,e.Hora,e.Texto)

}







