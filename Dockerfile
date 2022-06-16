FROM golang:1.17-alpine as builder

RUN apk add g++ && apk add make

# ENV 设置环境变量
ENV GOPATH=/opt/repo
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 输出目录
RUN mkdir /app

WORKDIR $GOPATH/go-seed-project

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . .
# 生成graphql文件
RUN make gen

# build main
RUN go build \
    -o "/app" \
    "main.go"

FROM alpine:latest

# 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# 安装依赖
RUN apk --no-cache add  \
    ca-certificates \
    libc6-compat \
    libstdc++ \
    file \
    tzdata \
    git \
    bash

RUN mkdir -p /opt/oswin

# 拷贝二进制文件
COPY --from=builder /app/* /opt/oswin/go-seed-project
RUN chmod +x /opt/oswin/go-seed-project
WORKDIR /opt/oswin

EXPOSE 8000/tcp

CMD ["/opt/oswin/go-seed-project"]


