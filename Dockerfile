#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

#Actualizamos e instalamos make.
RUN apt-get update && apt-get install make

#Creamos un usuario con el cual haremos la ejecucion sin privilegios.
RUN adduser -disabled-password juanalberto58
USER juanalberto58

#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test
COPY . .

#Comando que activarán los test del proyecto.
CMD ["make", "test"] 
