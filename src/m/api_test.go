package m

import (
	"testing"
)

var (
	entr Entrada
	entrP Entrada
)

/****************************** TEST CONSTRUCTORES ******************************/

func TestContructor (t *testing.T){
	t.Log("Test Contructor")
	entr.Contructor()
	if(entr.GetTitulo()!=""){
		t.Error("Fallo en el contructor")
	}
}

func TestContructorParametros(t *testing.T){
	t.Log("Test ContructorParámetros")
	entrP.ContructorParametros("TituloTest","DiaTest","HoraTest","TextoTest",1)
	if(entrP.GetTitulo()!="TituloTest"){
		t.Error("Fallo en el contructor de parámetros")
	}
}

/****************************** TEST TITULO ******************************/

func TestGetTitulo(t *testing.T){
	t.Log("Test getTitulo")
	if (entrP.GetTitulo() != "TituloTest"){
		t.Error("Fallo en la función getTitulo, el valor es... "+entrP.GetTitulo()+" y el valor que debería tener el campo es TituloTest")
	}
}

func TestSetTitulo(t *testing.T){
	t.Log("Test setTitulo")
	entrP.SetTitulo("TestTituloAprobado")
	if(entrP.GetTitulo()!="TestTituloAprobado"){
		t.Error("Fallo en la función setTitulo, el valor es... "+entrP.GetTitulo()+" y el valor que debería tener el campo es TestTituloAprobado")
	}
}

/****************************** TEST DIA ******************************/

func TestGetDia(t *testing.T){
	t.Log("Test getDia")
	if (entrP.GetDia() != "DiaTest"){
		t.Error("Fallo en la función getDia, el valor es... "+entrP.GetDia()+" y el valor que debería tener el campo es DiaTest")
	}
}

func TestSetDia(t *testing.T){
	t.Log("Test setDia")
	entrP.SetDia("TestDiaAprobado")
	if(entrP.GetDia()!="TestDiaAprobado"){
		t.Error("Fallo en la función setDia, el valor es... "+entrP.GetDia()+" y el valor que debería tener el campo es TestDiaAprobado")
	}
}

/****************************** TEST HORA ******************************/

func TestGetHora(t *testing.T){
	t.Log("Test getHora")
	if (entrP.GetHora() != "HoraTest"){
		t.Error("Fallo en la función getHora, el valor es... "+entrP.GetHora()+" y el valor que debería tener el campo es HoraTest")
	}
}

func TestSetHora(t *testing.T){
	t.Log("Test setHora")
	entrP.SetHora("TestHoraAprobado")
	if(entrP.GetHora()!="TestHoraAprobado"){
		t.Error("Fallo en la función setHora, el valor es... "+entrP.GetHora()+" y el valor que debería tener el campo es TestHoraAprobado")
	}
}

/****************************** TEST TEXTO ******************************/

func TestGetTexto(t *testing.T){
	t.Log("Test getTexto")
	if (entrP.GetTexto() != "TextoTest"){
		t.Error("Fallo en la función getTexto, el valor es... "+entrP.GetTexto()+" y el valor que debería tener el campo es TextoTest")
	}
}

func TestSetTexto(t *testing.T){
	t.Log("Test setTexto")
	entrP.SetTexto("TestTextoAprobado")
	if(entrP.GetTexto()!="TestTextoAprobado"){
		t.Error("Fallo en la función setTexto, el valor es... "+entrP.GetTexto()+" y el valor que debería tener el campo es TestTextoAprobado")
	}
}

/****************************** TEST CONT ******************************/

func TestGetCont(t *testing.T){
	t.Log("Test getCont")
	if (entrP.GetCont() != 1){
		t.Error("Fallo en la función getCont, el valor es... ",entrP.GetCont()," y el valor que debería tener el campo es 1")
	}
}

func TestSetCon(t *testing.T){
	t.Log("Test setCont")
	entrP.SetCont(1)
	if(entrP.GetCont()!=1){
		t.Error("Fallo en la función setCont el valor es... ",entrP.GetCont()," y el valor que debería tener el campo es 1")
	}
}

/****************************** TEST GUARDAR_ENTRADA ******************************/
func TestGuardarEntrada(t *testing.T){
	t.Log("Test GuardarEntrada")
	entr.GuardarEntrada("testTituloGe","testDiaGe","testHoraGe","testTextoGe")
	if(entr.GetTitulo()!="testTituloGe"){
		t.Error("Fallo en la función GuardarEntrada")
	}
}
