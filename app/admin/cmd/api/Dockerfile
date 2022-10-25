FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
RUN go mod download
COPY . .
COPY app/admin/cmd/api/etc  /app/etc

RUN go build -ldflags="-s -w" -o /app/core ./app/admin/cmd/api/core.go

FROM golang:1.19-alpine

RUN apk update --no-cache && apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/core /app/core
COPY --from=builder /app/etc /app/etc

CMD ["./core", "-f", "etc/core.yaml"]
