package main

type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
}

/****************************** CONSTRUCTORES ******************************/
func (e *Entrada)Contructor(){
	e.Titulo = ""
	e.Dia = ""
	e.Hora = ""
	e.Texto = ""
}

func (e *Entrada)ContructorParametros(tit string, d string, h string, text string){
	e.Titulo = tit
	e.Dia = d
	e.Hora = h
	e.Texto = text
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


/***************************** GUARDAR ENTRADA ******************************/
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