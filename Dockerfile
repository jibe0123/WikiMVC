FROM golang:1.14.1-alpine3.11 as builder
LABEL maintainer="Agostin Jean-baptiste <Jbagostin@gmail.com>"

ENV GO111MODULE=on

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/* && apk add ca-certificates

WORKDIR /go/WikiMVC
COPY go.sum go.mod /go/WikiMVC/

RUN go mod download

COPY . /go/WikiMVC/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/WikiMVC .

FROM alpine:latest

COPY --from=builder /go/WikiMVC/bin/WikiMVC /
ADD  web ./web

EXPOSE 8080

ENTRYPOINT ["/WikiMVC"]


