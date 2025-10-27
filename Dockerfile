# Build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary and assets
COPY --from=builder /app/main .
COPY --from=builder /app/web ./web
COPY --from=builder /app/assets ./assets

EXPOSE 8080

CMD ["./main"]
