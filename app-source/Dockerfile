# Build stage
FROM golang:latest AS builder

ENV GOOS=linux
ENV GOARCH=arm64
ENV GO111MODULE='on'

WORKDIR /src
COPY src .
RUN go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/main .

EXPOSE 8080
CMD ["./main"]

