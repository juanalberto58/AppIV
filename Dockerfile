#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

USER root 

#Actualizamos, instalamos make y además añadimos un usuario aprovechando la misma instrucción.
RUN apk update \
&& apk add make \
&& apk build-base \
&& apk add git \
&& go get -u github.com/gin-gonic/gin

RUN addgroup

#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test

#Comando que activarán los test del proyecto.
CMD ["make", "test"]
