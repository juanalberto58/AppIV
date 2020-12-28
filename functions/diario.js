const diario = require('./diario.json');

exports.handler = async function(event, context){
  if (diario.length > 0){
    var mensaje = "[";

	//Se leen todos los datos del diario
    for (var i = 0; i < diario.length; i++){

      mensaje += "{Titulo: " + diario[i]['Titulo'] + ", " +
        "Dia: " + diario[i]['Dia'] + ", " +
        "Hora: " + diario[i]['Hora'] + ", " +
        "Entrada: " + diario[i]['Entrada'] + "}";
    }

    mensaje += "]";
  }

  //Se muestran los datos
  if(diario!=""){
	 return{
	    statusCode: 200,
	    body: JSON.stringify(mensaje)
  	}
  }else{
  	return {
  		statusCode:500,
  		body:"No existe ninguna entrada que mostrar"
  	}
  }
}