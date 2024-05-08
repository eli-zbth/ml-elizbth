# Tech discovery

En el presente documento se puede encontrar toda la investigación previa que se realizó de forma previa a plantear la solución requerida.

## Contenido

- [Concepto de Url Shortener](##Concepto-de-URL-Shortener)
- [Arquitectura solución](##Arquitectura)


## Concepto de URL Shortener

Los Url Shortener son un tipo de servicio que permite crear una URL más pequeña a partir de una original la cual suele ser muy larga. De esta forma, el usuario puede interactura directamente con la nueva URL que es más manejable, facil de compartir y recordar. Dentro de las principales ventajas de usar este tipo de servicios se encuentra:

- Mejorar la Estética al usar cadenas de texto cortas e incluso personalizadas
- Permite insertar URL's en contextos con espacio reducido de caráctere, por ejemplo SMS o X(Twitter)
- Al ser fáciles de recordar, se vuelven útiles para incluir URL's en sitios web o publicaciones en formato físico.

Normalmente el dominio de la URL es actualizado a un nombre de dominio de nivel superior que sea más fácil de recordar y el conteido de esta se cambia por una cláve única que permite identificar la URL. En terminos de navegación, cuando el usuario accede a estas URL's acortadas los servicios suelen retornar respuestas 301 (redirección permanente) 302 o 307 (redirección temporal).

### Algoritmo de corte

Los algoritmos de corte se podrían clasificar en 3 tipos:

- Generación simple: La clave se compone de un id autoincremental guardado en una base de datos o también puede ser una cadena de carácteres aleatorios.
- Hash codificado: Se genera un hash mediante algoritos como MD5 o SHA-256 y posteriormente pasan algun tipo de encriptado (gemeralmente base32 o 64).
- Funciones biyectivas: Se utiliza una función que cumpla con la propiedad de ser biyectiva para generar la codificación de la url.

Los algoritmos de generación simple son los más fáciles de elaborar pero tienen la desventaja de tener la mayor probabilidad de generar una clave repetida. Por el contrario, los algoritmos basado en funciones biyectivas son los más seguros en terminos de que es infactible que genere una clave duplicada, sin embargo, solo existe 1 combinación posible para cada url, de modo que el algoritmo en si mismo no es capáz de cambiar y actualizar una clave. En base a lo anterior, los algoritmos de Hash codificado son los más utilizados cuando se crean este tipo de soluciones.


## Arquitectura

El diseño de solución más común para un URL shortener tiene el siguiente diagrama

![plot](./docs/img/generic-diagram.png)

- Una **interfaz web** que permite al usuario interactuar con el servicio
- **API** que al menos debe tener la funcionalidad de acortar una URL y poder redirigir a la url original cuando se ingresa la URL acortada.
- Una **base de datos** donde mapear la url original con su respectiva clave única acortada.
- Algún servicio de **caché** para agilizar el servicio de redirección.

referencias
https://www.karanpratapsingh.com/courses/system-design/url-shortener
https://www.milanjovanovic.tech/blog/how-to-build-a-url-shortener-with-dotnet
https://www.linkedin.com/pulse/system-design-development-url-shortener-service-1-2-syed-shah-kvr6e
https://medium.com/@sureshpodeti/system-design-url-shortener-d5d2e094a729
https://systemdesign.one/url-shortening-system-design/
https://www.geeksforgeeks.org/system-design-url-shortening-service/
https://medium.com/@sandeep4.verma/system-design-scalable-url-shortener-service-like-tinyurl-106f30f23a82