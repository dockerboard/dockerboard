FROM golang

MAINTAINER fundon cfddream@gmail.com

ADD . /go/src/github.com/dockerboard/dockerboard/
WORKDIR /go/src/github.com/dockerboard/dockerboard/

RUN go get ./...
RUN go build dockerboard.go

EXPOSE 8001

ENTRYPOINT ["./dockerboard"]
