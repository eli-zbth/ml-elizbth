# Ambiente cloud

En la siguiente carpeta se encuentra una serie de dockerfiles y configuraciones para levantar todos los elementos necesarios para hacer funcionar el sistea en formal local.

## ¿Por qué de esta forma?

Si bien existía al aternativa se solo crear un archivo dockercompose para levantar todas las instancias, preferí hacerlo de manera pensando en un ambiente productivo y bajo las siguientes razones:

- En un ecosistema productivo, no tiene sentido usar dockercompose porque cada artefacto es mantenido y deployado de forma independiente, incluso por equipos distintos.
- Sería una mala practica tener en un mismo repositorio la API y además el mantenimiento de toda su infraestructura.
- Al dejarlo en carpetas separadas se mantiene la api lo más "limpia" posible (en terminos de repo y deploy) simulando de mejor forma un ambiente productivo

## Instancias

### Mongo
Se debe ejecutar utiliando las siguientes instrucciones para así crear la base de datos con la colección e indices necesarias.

```bash

cd ./cloud/mongo  && docker build -t mongo_db_image . && cd ../..
docker run -d --name mongodb_container -p 27017:27017 mongo_db_image

```

Para acceder via consola se puede hacer de la siguiente forma

```bash

docker exec -it mongodb_container bash 
cd bin && mongosh
```


### Prometheus

Tambien se adjunta imagen docker con la customizacion necesaria para conectar prometheus a la API


```bash

cd ./cloud/prometheus  && docker build -t prometheus . && cd ../..
docker run -d --name prometheus_container -p 9090:9090 prometheus

```


### Grafana

Se dejó fuera del scoop hacer una configuración preestablecida para conectar automáticamente prometheus y guardar los dashboard, por ende se puede ejecutar con comandos estandar 

```bash

docker pull grafana/grafana
docker run -d --name my_grafana_container -p 3000:3000 grafana/grafana

```


### Redis

Ocurre caso similar que con grafana

```bash

docker pull redis
docker run -d --name my_redis_container -p 6379:6379 redis

```

### Api
Se inicia como cualquier proyecto de GO

```bash

go run main.go

```