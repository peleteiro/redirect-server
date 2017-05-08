FROM alpine:latest
MAINTAINER Jose Peleteiro <jose@peleteiro.net>

RUN apk update \
 && apk upgrade \
 && apk add \
            bash curl make git dumb-init \
            go libc-dev \
 && rm -rf /var/cache/apk/*

ADD . /tmp/build/src/github.com/peleteiro/redirect-server

RUN export GOPATH=/tmp/build \
 && go get golang.org/x/net/publicsuffix/... \
 && go build -o /usr/bin/redirect-server github.com/peleteiro/redirect-server \
 && rm -rf /tmp/build

ENV SERVER_FQDN redirect.foo.com
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/bin/redirect-server"]

EXPOSE 8080
