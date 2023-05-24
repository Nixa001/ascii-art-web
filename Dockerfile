FROM golang:1.20.4-alpine3.17 AS builder
RUN mkdir /app
LABEL version ="1.0" maintainer="@nifaye <nicolasfaye31@gmail.com>"
COPY . /app
WORKDIR /app
RUN apk update && apk add --no-cache bash
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080