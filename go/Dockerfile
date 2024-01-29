FROM golang:1.18-alpine as crudgo
WORKDIR /crudgo

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o crudgo crud.go

EXPOSE 8000
CMD ["./crudgo"]