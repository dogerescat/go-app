FROM golang:1.18-alpine

RUN apk update && apk add git

COPY ./src /app/
WORKDIR /app
RUN go mod tidy
#RUN go build -o main . 
RUN go build -o main

ENV GO111MODULE=on

EXPOSE 8080
#CMD ["go", "run", "main.go"]
CMD ["./main"]
