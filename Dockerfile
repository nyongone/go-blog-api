FROM golang:1.24-alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV DOCKERIZE_VERSION=v0.9.3

RUN apk update --no-cache \
    && apk add --no-cache wget openssl \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apk del wget

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o server cmd/server/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o migrate cmd/migrate/main.go

FROM alpine:3.21

COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /app /

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

CMD ["/server"]