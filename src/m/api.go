package m

type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
	cont int
}


/****************************** CONSTRUCTORES ******************************/
func (e *Entrada)Contructor(){
	e.Titulo = ""
	e.Dia = ""
	e.Hora = ""
	e.Texto = ""
	e.cont = 0
}

func (e *Entrada)ContructorParametros(tit string, d string, h string, text string, c int){
	e.Titulo = tit
	e.Dia = d
	e.Hora = h
	e.Texto = text
	e.cont = c
}
/****************************** GET and SET ******************************/

/**************** GET TITULO ***********************/
func (e Entrada)getTitulo() string{
	return e.Titulo
}

/**************** SET TITULO ***********************/
func (e *Entrada)setTitulo(tit string){
	e.Titulo=tit
}

/***************** GET DIA  ***********************/
func (e Entrada)getDia() string{
	return e.Dia
}

/***************** SET DIA  ***********************/
func (e *Entrada)setDia(d string){
	e.Dia=d
}

/***************** GET HORA ***********************/
func (e Entrada)getHora() string{
	return e.Hora
}

/***************** SET HORA ***********************/
func (e *Entrada)setHora(h string){
	e.Hora=h
}

/***************** GET TEXTO **********************/
func (e Entrada)getTexto() string{
	return e.Texto
}

/***************** SET TEXTO **********************/
func (e *Entrada)setTexto(t string){
	e.Texto=t
}

/***************** SET NUMERO DE ENTRADAS **********************/
func (e *Entrada)setCont(c int) {
	e.cont=c
}

/***************** GET NUMERO DE ENTRADAS **********************/
func (e Entrada)getCont() int{
	return e.cont
}

/***************** SEARCH ENTRADA **********************/
/**func searchEntrada(array []Entrada,d string, con int) (string,string,string){
	var i int
	var o int
	var encon bool
	var nada string

	for i = 0; i < con; i++ {
		if(array[i].getDia()==d){
			encon = true
		}else{
			encon = false
		}
	}
	if (encon==true){
		return array[i].getTitulo(), array[i].getHora(), array[i].getTexto()
	}else{
		return nada, nada, nada
	}

}*/


/****************************** GUARDAR ENTRADA ******************************/
/*func GuardarEntrada(tit string, d string, h string, text string){
	var e Entrada
	e.Titulo = tit
	e.Dia = d
	e.Hora = h
	e.Texto = text
}*/

/****************************** EDITAR ENTRADA ******************************/
/*func EditarEntrada(ntit string, nd string, nh string, ntext string){
	var e Entrada
	e.Titulo = ntit
	e.Dia = nd
	e.Hora = nh
	e.Texto = ntext
}*/
