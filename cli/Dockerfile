FROM golang:1.9.2

RUN mkdir -p /go/src \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

RUN go build -o cli .

CMD ["/go/src/app/cli"]