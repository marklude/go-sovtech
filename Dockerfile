# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV GO111MODULE=on

RUN mkdir /app
ADD . /app
WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build

EXPOSE 8081

CMD [ "./app" ]