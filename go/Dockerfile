FROM golang:1.18-alpine as crudgo
WORKDIR /crudgo

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o crudgo main.go

EXPOSE 8080
CMD ["./crudgo"]