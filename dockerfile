# Dockerfile
# 建立階段：使用 Go 官方映像進行建置
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 將 go.mod 和 go.sum 複製進容器並下載模組
COPY go.mod .
COPY go.sum .
RUN go mod download

# 複製所有原始碼並編譯
COPY . .
RUN go build -o slot-app

# 執行階段：使用輕量級映像來執行應用
FROM alpine:latest
WORKDIR /root/

# 從建置階段複製執行檔
COPY --from=builder /app/slot-app .

# 設定容器啟動時執行的指令
CMD ["./slot-app"]
