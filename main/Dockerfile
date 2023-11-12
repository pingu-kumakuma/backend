FROM golang:1.21.1 as build
WORKDIR /app
COPY . .
EXPOSE 8080
RUN go build main.go
CMD ["./main"]