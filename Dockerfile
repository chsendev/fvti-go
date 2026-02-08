FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.25.4-alpine AS builder
WORKDIR /data/code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o app

FROM scratch
COPY --from=builder /data/code/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]