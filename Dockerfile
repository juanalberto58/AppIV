#Utilizaremos la siguiente imagen de Golang, alpine.
FROM golang:alpine3.12

#Etiqueta de autor del dockerfile.
LABEL maintainer="Juan Alberto Rivera Peña"

#Actualizamos, instalamos make y además añadimos un usuario aprovechando la misma instrucción.
RUN apk add --no-cache make \
&& adduser --disabled-password juanalberto58 \
&& apk build-base \
&& apk add git \
&& go get -u github.com/gin-gonic/gin


#Asignamos directorio en el que trabajaremos y copiamos los archivos necesarios.
WORKDIR /test

#Comando que activarán los test del proyecto.
CMD ["make", "test"]
