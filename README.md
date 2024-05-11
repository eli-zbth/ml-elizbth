# Desafío ML

Hola!, mi nombre es Elizabeth Carreño y en este repositorio encontrarás mi propuesta de solución al [desafío de ML](./docs/challenge.md) correspondiente al proceso de selección.

## Aspectos principales del proyecto

- En [tech discovery](./docs/tech-discovery.md) se puede encontrar la documetación sobre la investigación previa que se realizó para tener contexto técnico y plantear la solución.

- Información sobre el [dimensionamiento de la plataforma](./docs/capacity-planning.md)

- Arquitectura propuesta y su explicación [arquitectura](./docs/architecture.md)


docker run -d --name prometheus -p 9090:9090 -p 10087:10087 -v  /Users/e0c02oi/Documents/personal/ml-elizbth/prometheus.yaml prom/prometheus --config.file=/etc/prometheus/prometheus.yml