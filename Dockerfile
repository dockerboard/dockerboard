FROM golang

MAINTAINER fundon cfddream@gmail.com

RUN go get github.com/dockerboard/dockerboard

WORKDIR /go/src/github.com/dockerboard/dockerboard/

RUN go get ./...
RUN go build dockerboard.go

RUN git clone --depth 1 https://github.com/dockerboard/bluewhale.git /bluewhale

EXPOSE 8001

ENTRYPOINT ["./dockerboard"]
