FROM golang:1.15.2-alpine

RUN apk update && apk add git

# COPY api/ /app/api

# WORKDIR /app/api

RUN go build -o main . 

ENV GO111MODULE=on

CMD ["go", "run", "main.go"]