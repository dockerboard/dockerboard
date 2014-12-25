FROM golang

MAINTAINER fundon cfddream@gmail.com

ADD . /go/src/github.com/dockerboard/dockerboard/
WORKDIR /go/src/github.com/dockerboard/dockerboard/

RUN git clone --depth 1 https://github.com/dockerboard/bluewhale.git /bluewhale
RUN go get ./...
RUN go build dockerboard.go

EXPOSE 8001

ENTRYPOINT ["./dockerboard"]
