# Dockerfile for UserService
FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service .
RUN chmod +x user-service
CMD ["/root/user-service"]
