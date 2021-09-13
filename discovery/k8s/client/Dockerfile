FROM golang:1.14-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/zero
COPY . .
RUN go build -ldflags="-s -w" -o /app/client discovery/k8s/client/client.go


FROM alpine

RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/client /app/client
COPY discovery/k8s/client/etc/k8s.yaml /app/

CMD ["./client", "-f", "k8s.yaml"]
