FROM golang:alpine AS build
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add build-base git
WORKDIR /src
RUN mkdir -p _output/linux
ARG TARGETARCH=amd64 VERSION=
COPY . .
RUN go build -tags "sqlite_foreign_keys release linux" -ldflags="-s -w -X 'main.Version=${VERSION}'" -o _output/linux/yarr src/main.go

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates
COPY --from=build /src/_output/linux/yarr /usr/local/bin/yarr
EXPOSE 7070
CMD ["/usr/local/bin/yarr", "-addr", "0.0.0.0:7070", "-db", "/data/yarr.db"]
