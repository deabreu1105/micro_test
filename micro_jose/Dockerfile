FROM golang:1.16-alpine
WORKDIR /usr/local/go/src/main
COPY . .
RUN go build -o main
EXPOSE 4000
CMD ["./main"]
