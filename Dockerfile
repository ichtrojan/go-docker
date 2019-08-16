FROM golang:latest

WORKDIR /

COPY . .

RUN go get -u github.com/gorilla/mux

CMD ["go", "run", "main.go"]
