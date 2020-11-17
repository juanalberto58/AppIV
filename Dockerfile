#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

#Actualizamos, instalamos make y además añadimos un usuario aprovechando la misma instrucción.
RUN apk update && apk add make && adduser -D juanalberto58

#Utilizaremos el usuario con el cual haremos la ejecucion sin privilegios.
USER juanalberto58

#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test
COPY . .

#Comando que activarán los test del proyecto.
CMD ["make", "test"] 
