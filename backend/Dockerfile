FROM golang:1.23.3-alpine

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 复制go mod文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 创建tokenList目录
RUN mkdir -p tokenList

# 构建应用
RUN go build -o main exec/main.go

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"] 