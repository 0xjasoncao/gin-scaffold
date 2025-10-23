FROM golang:1.23.5-alpine3.21 AS build-env

ARG VERSION

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod download && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -o /go/src/app/output/web /go/src/app

FROM alpine:3.21
COPY --from=build-env /go/src/app/output/web /web
COPY configs/dev /configs

CMD ["/web", "start", "-c", "/configs"]
