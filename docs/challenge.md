
# Descripción del desafío

Se necesita hacer un urlshortener tipo goo.gl o bitly para publicar promociones en twitter.

## Características esperadas

- Las urls deben tener vigencia indefinida
- La plataforma debe soportar un tráfico de alrededor 1M RPM.
- Debe contar con estadísticas de acceso nearly real time.
-Las url deben poder ser administradas:
  - Habilitar o desahibilitar
  - Modificar url de destino (cualquiera de sus componentes).
-La resolución de urls debe ser lo más rápido posible y con el menor costo.

## Entregables deseados

- Código fuente en repositorio privado de GitHub o Gitlab
- Documentación que indique cómo ejecutar el programa
- Documentación del proyecto que considere relevante
- URL en donde este hosteado el servicio(en caso de ser necesario)
Contemplar buenas prácticas

## Aspectos a consdierar

- Plantear distintos componentes explicando su responsabilidad y por que se incluyeron.
- Explicación de la infraestructura, herramientas/tecnologías preexistentes y cuales
son los motivos de la elección de cada una (por ejemplo si se incluye una ddbb en
particular, comentar si se evaluaron otras alternativas y cuál fue el racional de la
decisión final).
- Deseable incluir gráficos y explicación breve escrita para que la propuesta sea
correctamente entendida
- Deseable entregar una API Rest funcional corriendo en algún cloud público, con
una capacidad aproximada de 5000 rpm (verificable) a modo de demo (no hace falta
gastar dinero en la prueba, mockear lo que no pueda obtenerse gratis).
- El código debe ser compartido a través de un repositorio o bien en un zip.
