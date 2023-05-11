FROM golang:latest

WORKDIR /app

COPY . .

# RUN go get github.com/githubnemo/CompileDaemon

# ENTRYPOINT CompileDaemon -command="./app"

RUN go build -o main .

CMD ["./main"]