FROM alpine:latest
MAINTAINER Jose Peleteiro <jose@peleteiro.net>

RUN apk update \
 && apk upgrade \
 && apk add \
            s6 bash curl make git \
            go \
 && rm -rf /var/cache/apk/*

ADD ./root /

ADD . /tmp/build/src/github.com/peleteiro/redirect-server

RUN export GOPATH=/tmp/build \
 && go build -o /usr/bin/redirect-server github.com/peleteiro/redirect-server \
 && rm -rf /tmp/build

CMD ["/bin/s6-svscan", "/etc/s6"]

EXPOSE 8080
