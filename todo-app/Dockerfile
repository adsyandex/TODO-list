FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o todo-app cmd/todo-app/main.go

CMD ["./todo-app"]
