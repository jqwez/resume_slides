FROM golang:latest

WORKDIR /app

COPY src /app
RUN go build main.go

EXPOSE 80

CMD ["/app/main"]
