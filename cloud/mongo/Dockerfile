
FROM mongo:4.4


COPY mongo-init.js /docker-entrypoint-initdb.d/


EXPOSE 27017


CMD ["mongod", "--bind_ip_all"]