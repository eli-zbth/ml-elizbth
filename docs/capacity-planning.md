# Capacidad de la plataforma

Para fines de todos los cálculos anexados se considera que la plataforma recibirá **1 millón de request por minuto** y que cada registro de url puede pesar unos 600 bytes.

## Resumen Estimaciones

|Parámetro|  Valor|
|---|---|
|Request por segundo|  16.667 |
|Trafico por segundo| 1.000 Mb/seg |
|Almacenamiento para 5 años| 37 Tb|
|Largo optimo para url's acordadas| 8 carácteres|
|Nro de posibles urls| 218.240 Billones|
|Caché requerido| 173 Gb|


## Tráfico

Dada la volumetría de peticiones por minuto, se espera el siguiente tráfico de red para el servicio:

> *1.000.000 RPM ~  16.667 Request por segundo x 600 bytes = ***1000 MB/seg***

## Almacenamiento

 Se puede estimar una taza de crecimiento de 864 Gb de almacenamiento, lo que constituye un total de **4,32 Tb anuales**.

Proyectando el servicio a 5 años y con una taza de crecimiento del 10% respecto a las RPM se obtiene el siguiente resumen sobre el requerimiento de almacenamiento.

|Años|RPM|Demanda anual (TB)| Demanda Acumulada (TB)|
|-|-|-|-|
|1|1.000.000|4,32| 4,32|
|2|1.100.000|4,81|20.53|
|3|1.210.000|5,23|24.84|
|4|1.331.000|5,7|30.06|
|5|1.464.100|6,32|36.67|

## Estimación carácteres url acortada

Consdierando que se utiliza un algoritmo de base62 (es decir, posee 62 carácteres) para codificar las url acortadas, se debe encontrar un equilibrio entre la cantidad de carácteres y la variablidad que permite dicho largo para generar urls únicas. Se puede calcular el número posible de urls con la siguiente formula donde n es el largo de la cadena de carácteres

> *Number of URLs* = $62^n$

Si esto lo complementamos con la cantidad de request por minuto esperadas, podemos calcular la vida útil que nos puede proveer cada largo de caracter.

|n|Combinaciones de url posible|ejemplo| Vida útil esperada |
|-|----------------------------|-------|--------------------|
|5| 916 Mi                     | asdgt |          15,3 horas|
|6| 56.800 Mi                  | agtjuo| 39,4 días |
|7| 3.552 Bi | age4hui| 6,7 año |
|8| 218.240 Bi| efyjfghj| 415 años |
|9| 13.537 Tri| sffghjjrf| 25.755 años |
|10|  839.299  Tri| dghrfhyjyf| 1.596.840 años|

Considerando lo anterior, una cadena de 8 carácteres en condificada en base62 provee la cantidad de combinaciones necesarias para un proyecto de duración practicamente indefinida, siendo así un método seguro para generar claves únicas. Otra forma de confimar esto es revisando cuanto cambia la vida útil del servicio si aumenta la cantidad de request por minuto

|Request por minuto| Vida útil esperada |
|------------------|---------------------|
|1.000.000| 415 años |
|2.000.000| 208 años |
|5.000.000| 83 años |
|16.000.000| 28 años |

De esta forma, se puede confirmar que una url acortada de 8 carácteres es suficiente para asegurar integridad de las claves generadas.


## Demanda de caché

Para calcular el caché requerido se respetará la regla del 80/20, en donde se espera que el 20% de las urls sean las responsables del 80% del tráfico total. De esta manera tenemos que tener la capacidad de almacenar 288.000.000 request equivalentes a 173 GB de caché diarios.

> *1.000.000 RPM al 20% => 200.000 RPM  ~ 288.000.000 request por día*
> *288.000.000 * 500 bytes = 173 GB*
