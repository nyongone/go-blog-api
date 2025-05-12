FROM golang:1.24-alpine as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o server cmd/server/main.go

FROM scratch

COPY --from=builder /app /

CMD ["/server"]