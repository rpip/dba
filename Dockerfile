FROM alpine:latest
MAINTAINER Yao Adzaku <yao.adzaku@gmail.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go
ENV GO15VENDOREXPERIMENT 1

RUN	apk add --no-cache \
	ca-certificates

COPY . /go/src/github.com/rpip/dba
COPY test-fixtures/sample.conf.hcl /usr/local/etc/dba.conf

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		go \
		git \
		gcc \
		libc-dev \
		libgcc \
	&& cd /go/src/github.com/rpip/dba \
	&& go get github.com/Masterminds/glide github.com/k0kubun/pp \
	&& glide install \
	&& go build -o /usr/bin/dba . \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

ENTRYPOINT ["dba"]

CMD ["--help"]
