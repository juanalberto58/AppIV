# Serverless

## Netlify

Para empezar lo que tenemos que hacer es ir a la página de [Netlify](https://www.netlify.com/), una vez dentro vamos a registrarnos y trás darle a registro ya nos aparace la opción de vincular con nuestra cuenta de Github.

![Netlify1](../image/netlify1.png)

Autorizamos el acceso:

![Netlify2](../image/netlify2.png)

Una vez hecha nuestra cuenta de Netlify ya vinculada con Github, el siguiente paso será crear nuestro sitio en Netlify que deberá de estar enlazado con nuestro repositorio de Github.

![Netlify3](../image/netlify3.png)

Tenemos que volver a autorizar el acceso y ya nos saltará el proceso de instalación de Netlify en el que tendremos que seleccionar nuestro repositorio de Github y posteriormente instalar: 

![Netlify4](../image/netlify4.png)

Una vez hecho todos los pasos anteriores lo que nos queda es configurar el despliegue:

![Netlify5](../image/netlify5.png)

Tras realizar todos los anteriores pasos ya tenemos Netlify configurado y nos proporciona un dominio por defecto, el cual vamos a cambiar por **appiv.netlify.app**

![Netlify6](../image/netlify6.png)
![Netlify7](../image/netlify7.png)

Vamos a configurar tambien los despliegues para ello que vamos a decir a Netlify que el directorio en el cual encontrará nuestros ficheros de código será **functions**:

![Netlify8](../image/netlify8.png)

Una vez todo configurado lo siguiente es hacer la función que se implementará, dicha [función](https://github.com/juanalberto58/AppIV/blob/master/functions/diario.js) pertenece a la [historia de usuario](https://github.com/juanalberto58/AppIV/issues/32).
Como podemos ver dicha función lo que realiza es que muestra todas las entradas registradas en el diario y estas entradas se encuentran en el fichero diario.json las cuales las devuelve si todo se ha realizado correctamente.

A través del siguiente [enlace](https://appiv.netlify.app/.netlify/functions/diario) se puede ver como dicha función funciona correctamente. 
