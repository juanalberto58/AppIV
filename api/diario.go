package handler

import (
    "fmt"
    "net/http"
    "time"
    "github.com/juanalberto58/AppIV/src/YourLife"
  	"encoding/json"
  	"strings"
)

var diario = []Entrada {
	Entrada {
		titulo: "Me voy de pesca",
		dia: "Lunes",
		hora: "17:00",
		entrada: "Hoy me he ido de pesca al rio"
	},
	Entrada {
		titulo: "Me voy de caza",
		dia: "Martes",
		hora: "11:00",
		entrada: "Hoy me he ido de caza al monte"
	},
	Entrada {
		titulo: "Me voy a comer a un restaurante",
		dia: "Miercoles",
		hora: "14:00",
		entrada: "Hoy me he ido a comer a un restaurante"
	},
	Entrada {
		titulo: "Me voy a Sevilla",
		dia: "Jueves",
		hora: "10:00",
		entrada: "Hoy me he ido a Sevilla"
	},
	Entrada {
		titulo: "Me voy a Jaen",
		dia: "Viernes",
		hora: "11:00",
		entrada: "Hoy me he ido a Jaen"
	},

}


func Handler(w http.ResponseWriter, r *http.Request) {
	var tope int
	tope = range diario

	for i, Entrada := range diario {
		if ( i <= tope ) {
			fmt.Fprintf(w, Entrada.getTitulo())
			fmt.Fprintf(w, Entrada.getEntrada())
		} else {
			fmt.Fprintf(w, "No existen entradas" )
		}
	}
	
}