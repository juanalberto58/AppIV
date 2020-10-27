# Herramientas

## Lenguaje **Go**

Para este proyecto he elegido el lenguaje de programación Go ya buscando información sobre este, he visto que es relativamente sencillo de programar y nos ofrece una cantidad de herramientas considerable que nos abre un sinfin de posibilidades. Además por la sencillez de integración con la base de datos MongoDb y que nos proporciona los paquetes necesarios para realizar test como GoTest sin tener que utilizar librerías externas.

## Base de Datos **MongoDb**

Como base de datos he escogido MongoDb ya que ya he trabajado en otras ocasiones con el y lo conozco además de por la potencia que ofrece y la naturalidad de concebir los datos frente al modelo tradicional de filas y columnas.


## Test **GoTest**

Para realizar los test el propio lenguaje Go ya nos provee de una entorno de pruebas sin tener que recurrir a bibliotecas externas. 

Esta biblioteca propia de test a diferencia de otros lenguajes de programación, no utiliza aserciones, utiliza una serie de estructuras de datos que al utilizarlas nos informan de los errores que se hayan podido producir.

De modo que para este proyecto al elegir Go como lenguaje no tenemos que recurrir a utilidades externas ya que dicho lenguaje ya nos las proporciona.


## Contrucción **Makefile**

El lenguaje Go tiene un gestor de tareas implícito en el compilador y además aparentemente no existe ningún task runner específico para este lenguaje, por lo tanto la herramienta de construcción escogida es [Makefile](https://github.com/juanalberto58/AppIV/blob/master/makefile) ya que ya lo he utilizado previamente y lo conozco bastante bien. Makefile nos ofrece además mucha facilidad a la hora de poder tener todas las tareas juntas y poder ejecutarlas de manera automática. En este caso ahora mismo las tareas que podremos llevar a cabo son las siguientes:
	- make build: Compilamos el proyecto.
	- make run: Ejecutamos el proyecto.
	- make test: Realizamos los test al proyecto.







