#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

ENV CGO_ENABLED=0

#Actualizamos, instalamos make y además añadimos un usuario aprovechando la misma instrucción.
RUN apk add --no-cache make \
&& apk add build-base 

#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test

#Comando que activarán los test del proyecto.
CMD ["make", "test"]
