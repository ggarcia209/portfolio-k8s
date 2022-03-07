# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./

COPY ./go.sum ./

COPY ./index.html ./

RUN go mod download

COPY *.go ./

RUN go build -o .

EXPOSE 8081

CMD ["./serve"]