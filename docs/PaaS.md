# PaaS

## Descripción y justificación para desplegar en un PaaS

La plataforma que usaremos será Heroku, ¿Por que Heroku?, lo primero es porque ya lo he utilizado en otras ocasiones, lo conozco y se cuales son las funcionalidades que nos ofrece, además de soportar aplicaciones en el lenguaje Go.

Heroku nos ofrece una documentación excelente y su utilización para arrancar nuestra aplicación es sencilla además de ofecernos la posibilidad de ser gratuito el proceso. En Heroku podemos instalar addons para utilizar funcionalidades nuevas, tiene un sistema de log bastante bueno, sobre todo nos permite escalar nuestra aplicación si fuera necesario.

## Descripción de la configuración

Lo primero que hay que hacer es registrarse y abrir una cuenta en Heroku, una vez hemos creado nuestra cuenta lo siguiente es ir a crear un nueva app, añadimos el nombre que queramos que tenga nuestra app, en mi caso 'appivheroku' y Europa de región.

Lo siguiente que vamos a hacer es irnos al menú Deploy y vamos a vincular nuestra app de heruko con nuestro repo de github:

![heroku1](../image/heroku1.png)

Buscamos nuestro repo y lo conectamos, una vez hecho esto ya tenemos nuestro repo conectado con nuestra app de heroku:

![heroku2](../image/heroku2.png)

Una vez hecho esto, el siguiente paso es realizar el archivo de configuración de Heroku en nuestra aplicación para que Heroku sepa que acción tiene que realizar. Este archivo se llama Procfile y lo utiliza Heroku para arrancar la aplicación, en mi casa en este archivo habrá una llamada a una acción en mi task-runner que iniciara el proceso.

