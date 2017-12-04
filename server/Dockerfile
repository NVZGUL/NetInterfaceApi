FROM golang:1.9.2

RUN mkdir -p /go/src \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app
RUN  go get github.com/gorilla/mux

RUN go build -o server .

CMD ["/go/src/app/server"]
