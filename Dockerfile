FROM golang:1.23
LABEL authors="miguel soto"

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

# Comando para iniciar la aplicación
CMD ["./main"]
