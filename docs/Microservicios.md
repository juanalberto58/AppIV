# Microservicios

## Justificación de Framework Gin-Gonic

El Microframework para la implementación del microservicio elegido es [Gin-Gonic](https://github.com/gin-gonic/gin) ya que tras hacer un estudio de las diferentes posibilidades como por ejemplo [Beego](https://beego.me/), [Martini](https://github.com/go-martini/martini) o [Revel](https://revel.github.io/), Gin nos ofrece un entorno minimalista y solo nos incluye las bibliotecas y funcionalidades esenciales necesarias para nuestro proyecto, si nuestro proyecto fuera de mayor envergadura la elección de microframework sería Revel ya que nos ofrece una cantidad enorme de funciones y posibilidades. 

Además Gin nos ofrece una muy buena documentación y sobretodo la gran virtud de este microframework es su velocidad, ya que al estar pensado para ser lo mas minimalista posible es extremadamente rápido en comparación con por ejemplo Revel o Martini, siendo hasta 40 veces más rápido que este último. 

A través de este [enlace](https://github.com/gin-gonic/gin#benchmarks) podrás ver los Benchmark de algunos de los Microframeworks citados anterioremente.


## Diseño general de la api y Historias de Usuario

En la Api se han desarrollado las siguientes HU en el archivo **main.go** en el cual se utiliza el archivo **api.go** el cual tiene el código necesario para realizar las funciones internas de manera que realicen las distintas acciones, además **main.go** tiene sus respectivos test en el archivo **main_test** que los ejecuto importando el paquete httptest de GO que nos proporciona todas las herramientas necesarias para realizar los test de servidor diciendonos si las rutas y lo que estas ejecutan al recibir la petición nos devuelve el resultado correcto. La **api.go** también tiene sus respectivos test pero esta vez solo importamos el paquete **testing** de Go con el cual realizamos los test automatizados.

Cada HU implementada tiene su respectiva ruta diferente dentro de la app, las HU implementadas son las siguientes:

- HU1: El usuario podrá añadir una nueva entrada al diario.
	- Esta HU se desarrolla con la función introducirEntrada(), en ella se procesan y se envian los datos de una entrada mediante una petición POST y su ruta es /anadeEntrada.
- HU2: El usuario podra editar la entrada del diario.
	- Una vez introducida la entrada mediante la función editarEntrada() podemos editar la entrada enviada. Además esta se realiza también con una petición POST y su ruta es /modificarEntrada.
- HU4: El usuario podrá ver el número de entradas del diario.
	- Esta HU se implementa en la función obtenerNumEntradas(), esta se obtiene mediante una petición GET y su ruta es /numeroEntradas.
- HU6: El usuario podrá consultar una entrada del diario.
	- Se implementa en obtenerEntrada() la cual obtenemos con una petición GET y su ruta es /obtenerEntrada.

## Uso de buenas practicas, middleware y log
