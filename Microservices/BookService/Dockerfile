# Dockerfile for BookService
FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/book-service .
RUN chmod +x book-service
CMD ["/root/book-service"]
