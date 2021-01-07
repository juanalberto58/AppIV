package handler

import (


    "fmt"
    //"ap/src/main"
    "github.com/juanalberto58/AppIV/src/YourLife"
    "net/http"

)

/*type Entrada struct{
	Titulo string
	Dia string
	Hora string
	Texto string
}*/


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
	
	if r.URL.String() == "api/diario" {
			fmt.Fprintf(w, "Funciona")
	}else{
		fmt.Fprintf(w, "Funciona")
	}
}