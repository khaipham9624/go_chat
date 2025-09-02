# Stage 1: Build binary
FROM golang:1.25 AS builder

WORKDIR /app

# Copy go.mod và go.sum trước để cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary (CGO_ENABLED=0 để tạo binary tĩnh, portable hơn)
RUN CGO_ENABLED=0 GOOS=linux go build -o go_chat ./src/server

# Stage 2: Deploy binary
FROM alpine:latest

WORKDIR /app

# Copy binary từ stage builder
COPY --from=builder /app/go_chat .

# Nếu có static/template thì copy thêm
COPY --from=builder /app/src/server/static ./static
COPY --from=builder /app/src/server/templates ./templates

EXPOSE 8080
CMD ["./go_chat"]
