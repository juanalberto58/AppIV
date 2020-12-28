const diario = require('./diario.json');

exports.handler = async function(event, context){
  if (diario.length > 0){
    var mensaje = "[";

	//Se leen todos los datos del diario
    for (var i = 0; i < diario.length; i++){

      mensaje += "{titulo: " + diario[i]['titulo'] + ", " +
        "dia: " + diario[i]['dia'] + ", " +
        "hora: " + diario[i]['hora'] + ", " +
        "entrada: " + diario[i]['entrada'] + "}";
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