FROM golang:1.19.2-alpine3.16
RUN apk update && apk add git 
WORKDIR /go/src/app
COPY ./api /go/src/app
RUN go mod tidy
# CMD ["go", "run", "main.go"]