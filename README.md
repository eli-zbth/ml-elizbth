# Desafío ML

Hola!, mi nombre es Elizabeth Carreño y en este repositorio encontrarás mi propuesta de solución al [desafío de ML](./docs/challenge.md) correspondiente al proceso de selección.

## Aspectos principales del proyecto

- En [tech discovery](./docs/tech-discovery.md) se puede encontrar la documetación sobre la investigación previa que se realizó para tener contexto técnico y plantear la solución.

- Información sobre el [dimensionamiento de la plataforma](./docs/capacity-planning.md)

- Arquitectura propuesta y su explicación [arquitectura](./docs/architecture.md)

- El proyecto se encuentra desplegado en [railway](https://railway.app/) a modo de demo. Se está utilizando una cuenta gratuita con recursos limitados pero en caso de ser necesario puedo extender el trial o aumentar el plan para una prueba específica.

## Fuera de scoope

A continuación se dejan una serie de elementos que quedaron fuera de scoop pero que deberian ser incluidos en caso de un proyecto real

- Aumentar el coverage del proyecto en general
- Pruebas de integración
- El pipeline para que ejecute linter y test automáticamente
- Trabajo mediante pull request y no directo a main
- Proteger main
- Integrar una heramienta que maneje secretos como Vault o akeyless
- Mejor seguridad en el manejo de accesos a la base de datos
- Interfaz web.
- Manejo de ambientes de dev/stage

## Descripción

La Api consta de 5 enpoints diferentes clasificados en 3 tipos

### Crear

#### /create

Permite obtener una URL más pequeña a partir de una URL de cualquier sitio web. El id de esta url puede ser aleatorio o un valor entregado por el usuario.

```json
Ejemplo:
{
    "Url": "http://www.someweb.cl/some_very_long_path", 
    "custom_url": "tiny_path"

}
```

La api retornará la url lista para ser usada en el navegador

```json
output:
{
    "short_url": "http://www.localhost:8080/yaVvkxh"
}
```

### Editar

permiten cambiar cualquiera de los elementos de las url

#### /edit/short_url

Entregando la url acortada permite personalizarla y cambiar su nombre
```json
Ejemplo:
{
    "short_url": "https://www.localhost.com/YpbBUDM",
    "new_value" : "prueba_edit_key"
}
```
Si la url resuleve de forma exitosa entregará un mensaje de confirmación.


#### /edit/redirect_url

Entregando la url acortada cambiar la url original a la cual estaba asocida
```json
Ejemplo:
{
    "short_url": "https://www.shorturl.com/tiny",
    "redirect_url" : "https://www.youtube.com/watch?v=mCdA4bJAGGk"
}
```
Si la url resuleve de forma exitosa entregará un mensaje de confirmación y ahora la url en caso de ser utilizada en un navegador redireccionará a la nueva url entregada

#### /edit/url_status

Inhabilita o activa una url. Solo aquellas que esten "activas" son redireccionadas.
```json
Ejemplo:
{
    "shorturl": "https://www.shorturlcom/d7X623s",
    "isactive": false
}
```
Si la url resuleve de forma exitosa entregará un mensaje de confirmación.


### Redirección

En este caso, se puede acceder al navegador directamente con la URL entregada por el endpoint de Create. Esta petición llega a la API y genera una redirección automática en caso de que la url entregada exista en la plataforma y esté activa.

### Otros

### /health

La api cuenta con un healthcheck básico

### /docs/swagger/
Se dejó activo para ambiente productivo el swagger de la Api a modo de tener documentación más clara. Además en la sección docs/postman se encuentra un postman disponible

## Ambiente local

En la sección [cloud](./cloud/cloud.MD) se dejaron todas las intrucciónes para levantar los artefactos necesarios para levantar el proyecto en local

### Desarrollo

#### Dependencias
1. Para el proyecto se necesita tener instalado *Go 1.22
2. Además se recomienta instalar *golangci-lint* y *swagger*

3. Instalación de packages
```
    - Run `go mod tidy`
    - Run `go mod vendor`
```
#### Ejecución
Variables de entorno necesarias en archivo .env en carpeta raiz
```
    PORT=8080
    MONGO_DB_URI="mongodb://localhost:27017"
    SHORT_URL_DOMAIN="http://www.localhost:8080/"
    REDIS_URL="localhost:6379"
```


```
    go run main.go
```

#### Test
```
   go test ./...
```

#### Test y Coverage
```
    go test ./... -coverprofile=coverage.out   
    go tool cover -html=coverage.out
```