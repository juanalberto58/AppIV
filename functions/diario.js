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

  //Si no hay datos se muestra dicho mensaje
  else{
    mensaje = "{No existen entradas en el diario}";
  }

  //Se muestran los datos
  return{
    statusCode: 200,
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(mensaje)
  };
}