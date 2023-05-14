    FROM golang:latest
    # LABEL maintainer="nifaye <nicolasfaye31@gmail.com>"
    WORKDIR /app
    COPY . .
    RUN go build -o main .
    EXPOSE 8080
    CMD ["./main"]