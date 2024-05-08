# Estimación plataforma

## Tráfico

Por requisito de la solución sabemos que esta debe ser capáz de soportar **1 millón de request por minuto**

## Almacenamiento y vigencia de URL

Según los requisitos las url deben tener una duración indefinida, por lo que para fines de cálculos de almacenamiento se interpretará como un N grande y tendremos distintos escenarios

|  Atributo| valor |  
|----------|-------|
|  Vida útil url | 100 años   |
|  Request por minuto | 1 M por minuto =>  525.600 M request por año  |
|  Peso por registro | 500 bytes |

Considerando esa cantidad de request por segundo, la cantidad de información crecerá a tazas de 720 Gb por día, llegando a una capacidad máxima de poder soportar 52.560 Bi de registros equivalentes a 26.280 TB de memoria.

Cabe mencionar que tener válida una URL por 100 años es un número bastante grande y puede cambiar mucho la necesidad de almacenamiento en función de dicho valor.

|  Vida util Url | Almacenamiento requerido en capacidad máxima (TB) |  
|----------|-------|
|  100 |   26.280 |
|  50  |  13.140  |
|  25  |  6.570 |
|  20  |  5.256|

Podría considerarse que 25 años es un tiempo prudente para mantener viva la URL y baja considerablemente los requisitos de memoria para manternlas activas, permitiendo así que la plataforma pueda escalar en otra dimensiones (por ejemplo, solventar una mayor cantidad de request sin desbordar el requerimiento de memoria)