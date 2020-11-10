#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:1.15.3-alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test
COPY makefile .

#Actualizamos e instalamos make.
RUN apk update && apk add make

#Creamos un usuario con el cual haremos la ejecucion sin privilegios.
RUN adduser -D juanalberto58
USER juanalberto58

#Comando que activarán los test del proyecto.
CMD ["make","test"]
