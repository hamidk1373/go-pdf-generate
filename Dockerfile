FROM golang:1.21.3

RUN apt update

RUN apt upgrade -y

RUN apt install wkhtmltopdf -y

WORKDIR /app


# RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
RUN go mod download

CMD air
