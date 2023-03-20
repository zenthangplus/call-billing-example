FROM golang:1.18-alpine AS builder
ARG BUILD_MODULE
ARG BUILD_VERSION
ARG BUILD_COMMIT_HASH
ARG BUILD_TIME
ARG BS_PKG=github.com/zenthangplus/call-billing-example/src/$BUILD_MODULE/bootstrap
ENV GO111MODULE=on
RUN apk add --no-cache git
COPY . /go/app
WORKDIR /go/app
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./binary \
    -ldflags="-X '$BS_PKG.Version=$BUILD_VERSION' -X '$BS_PKG.CommitHash=$BUILD_COMMIT_HASH' -X '$BS_PKG.BuildTime=$BUILD_TIME'" \
    ./src/$BUILD_MODULE

FROM scratch
ARG BUILD_MODULE
COPY --from=builder /go/app/src/$BUILD_MODULE/config /app/config
COPY --from=builder /go/app/binary /app
EXPOSE 8080
WORKDIR /app
ENTRYPOINT ["./binary"]
