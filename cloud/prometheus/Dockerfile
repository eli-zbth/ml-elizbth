FROM prom/prometheus


COPY prometheus.yml /etc/prometheus/prometheus.yml

EXPOSE     9090

ENTRYPOINT [ "prometheus", \
             "--config.file=/etc/prometheus/prometheus.yml", \
             "--storage.tsdb.path=/prometheus", \
             "--storage.tsdb.retention=365d", \
             "--web.console.libraries=/usr/share/prometheus/console_libraries", \
             "--web.console.templates=/usr/share/prometheus/consoles", \
             "--log.level=info", \
             "--log.format=json" ] 
