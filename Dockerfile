FROM golang as build

ADD . /usr/local/go/src/singo

WORKDIR /usr/local/go/src/singo

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server
# 交叉编译
# RUN SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build -o api_server

# 二段构建
FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/singo/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN  chmod +x /usr/bin/api_server

ENTRYPOINT [ "api_server" ]
