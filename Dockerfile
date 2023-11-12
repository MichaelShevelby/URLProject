FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

#Выкачиваем завивисимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#Билдим
RUN go build -o app app/cmd/main.go

CMD ["app"]