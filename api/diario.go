package handler

import (
    "fmt"
    //"github.com/juanalberto58/AppIV/src/main"
    "net/http"
    //"../src/main"
)

type Entrada struct{
	Titulo string
	Dia string
	Hora string
	Texto string
}

func (e Entrada)getDia() string{
	return e.Dia
}

func (e Entrada)getTitulo() string{
	return e.Titulo
}


var diario = []Entrada {
	Entrada {
		Titulo: "Me voy de pesca",
		Dia: "Lunes",
		Hora: "17:00",
		Texto: "Hoy me he ido de pesca al rio",
	},
	Entrada {
		Titulo: "Me voy de caza",
		Dia: "Martes",
		Hora: "11:00",
		Texto: "Hoy me he ido de caza al monte",
	},
	Entrada {
		Titulo: "Me voy a comer a un restaurante",
		Dia: "Miercoles",
		Hora: "14:00",
		Texto: "Hoy me he ido a comer a un restaurante",
	},
	Entrada {
		Titulo: "Me voy a Sevilla",
		Dia: "Jueves",
		Hora: "10:00",
		Texto: "Hoy me he ido a Sevilla",
	},
	Entrada {
		Titulo: "Me voy a Jaen",
		Dia: "Viernes",
		Hora: "11:00",
		Texto: "Hoy me he ido a Jaen",
	},

}


func Handler(w http.ResponseWriter, r *http.Request) {
var encontrado bool
encontrado=false
var no_exis string 
no_exis=""

	if r.URL.String() == "api/diario" {
			var i int
			for i = 0; i < len(diario); i++ {
				if "dia" == diario[i].getDia(){
					w.Header().Add("Content-Type", "text/html")
					fmt.Fprintf(w, `
						<!DOCTYPE html>
							<head>
								<tittle>Titulo</tittle>
							</head>	
							<body>
								<h1>`+diario[i].getTitulo()+`</h1>
							</body>
					`)
				}
			}
		
	}else{
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<head>
				<tittle>Diario</tittle>
			</head>	
			<body>
				<form method="GET">
					<label>
						<input type="text" name="dia"/>
					</label>
					<button type="submit">Buscar</button>
				</form>
			</body>
			`)

			if r.Method == "GET" {
				var d = r.FormValue("dia")

				var i int
				for i = 0; i < len(diario); i++ {
					if d == diario[i].getDia(){
						encontrado=true
						w.Header().Add("Content-Type", "text/html")
						fmt.Fprintf(w, `
							<!DOCTYPE html>
								<body>
									<h1> Titulo de tu dia </h1>
									<h3>`+diario[i].getTitulo()+`</h3>
								</body>
						`)
					}else{
						no_exis="No existe ese dia"
					}
				}
				if encontrado == false {
					w.Header().Add("Content-Type", "text/html")
					fmt.Fprintf(w, `
						<!DOCTYPE html>
							<body>
								<h1> `+no_exis+`</h1>
							</body>
					`)
				}
					
			}

	}
}