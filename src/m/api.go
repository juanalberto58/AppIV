package m

type Entrada struct{
	Titulo string
	Dia string 
	Hora string
	Texto string
	cont int
}

/*type Entradas struct{
	array []Entrada
}*/


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


/*func NewEntradas() *Entradas{
	entradas := new(Entradas)
	return entradas
}*/

/**************** GET TITULO ***********************/
func (e Entrada)GetTitulo() string{
	return e.Titulo
}

/**************** SET TITULO ***********************/
func (e *Entrada)SetTitulo(tit string){
	e.Titulo=tit
}

/***************** GET DIA  ***********************/
func (e Entrada)GetDia() string{
	return e.Dia
}

/***************** SET DIA  ***********************/
func (e *Entrada)SetDia(d string){
	e.Dia=d
}

/***************** GET HORA ***********************/
func (e Entrada)GetHora() string{
	return e.Hora
}

/***************** SET HORA ***********************/
func (e *Entrada)SetHora(h string){
	e.Hora=h
}

/***************** GET TEXTO **********************/
func (e Entrada)GetTexto() string{
	return e.Texto
}

/***************** SET TEXTO **********************/
func (e *Entrada)SetTexto(t string){
	e.Texto=t
}

/***************** SET NUMERO DE ENTRADAS **********************/
func (e *Entrada)SetCont(c int) {
	e.cont=c
}

/***************** GET NUMERO DE ENTRADAS **********************/
func (e Entrada)GetCont() int{
	return e.cont
}


/****************************** GUARDAR ENTRADA ******************************/
func (e *Entrada)GuardarEntrada(tit string, d string, h string, text string){

	e.Titulo = tit
	e.Dia = d
	e.Hora = h
	e.Texto = text

	num_entr := e.GetCont()
	num_entr++
	e.SetCont(num_entr)
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

/****************************** EDITAR ENTRADA ******************************/
func (e *Entrada)EditarEntrada(ntit string, nd string, nh string, ntext string){
	e.Titulo = ntit
	e.Dia = nd
	e.Hora = nh
	e.Texto = ntext
}
