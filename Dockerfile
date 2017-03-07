FROM golang:latest

RUN mkdir -p /go/src/github.com/rpip/dba
WORKDIR /go/src/github.com/rpip/dba
ADD . /go/src/github.com/rpip/dba

RUN go get github.com/Masterminds/glide github.com/k0kubun/pp
RUN make deps && make build

ENTRYPOINT ["./dba"]

CMD ["--help"]
