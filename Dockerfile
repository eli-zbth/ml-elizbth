
# Initial stage: download modules
FROM golang:1.22.3 as builder

ENV config=docker

WORKDIR /app

COPY . /app

RUN go mod download


RUN go build -o /app

CMD [ "/app" ]