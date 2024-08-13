# 使用 golang:1.22 作为构建阶段的基础镜像
FROM golang:1.22 AS builder

# 设置环境变量
ENV GOPROXY='https://goproxy.io' \
    GOSUMDB='off' \
    GOOS='linux' \
    GOARCH='amd64' \
    CGO_ENABLED=0

# 创建一个 /app 目录来存放应用程序的源文件
RUN mkdir /app

# 将当前目录的所有文件复制到镜像中的 /app 目录
ADD . /app

# 设置工作目录为 /app
WORKDIR /app

# 编译 Go 程序，生成二进制文件 main
RUN go build -o paper-manager .

# 使用 alipine 作为基础镜像
FROM alpine:3

# 设置工作目录为 /app
WORKDIR /app

# 从构建阶段复制生成的二进制文件到当前镜像中的 /app 目录
COPY --from=builder /app/paper-manager /app

# 复制配置文件到镜像中的 /app/conf 目录
COPY ./conf/*.toml /app/conf/

# 设置一个构建参数 DEPLOY_ENV 并将其值设置为 develop
ARG DEPLOY_ENV=develop

# 设置环境变量 DEPLOY_ENV
ENV DEPLOY_ENV=$DEPLOY_ENV  

# 暴露端口 8090
EXPOSE 8090

# 设置容器启动时执行的命令，启动 Go 应用程序
ENTRYPOINT ["/app/paper-manager"]