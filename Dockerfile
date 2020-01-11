# Builder
FROM golang:alpine as builder

WORKDIR /go/src/manslaughter03/webmail

ENV GO111MODULE=on

ADD go.mod .
ADD go.sum .

RUN apk add --update git \
  && go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o app .

# Final image
FROM alpine:latest

RUN apk add --no-cache --update ca-certificates tzdata \
  && cp /usr/share/zoneinfo/Europe/Paris /etc/localtime \
  && echo "Europe/Paris" > /etc/timezone \
  && mkdir -p /usr/src/app/data \
  && adduser -D webmail \
  && chown -R webmail:webmail /usr/src/app \
  && apk del tzdata

USER webmail
WORKDIR /usr/src/app

COPY --from=builder /go/src/manslaughter03/webmail .

ENTRYPOINT ["./app"]

