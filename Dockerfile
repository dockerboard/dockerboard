FROM golang

MAINTAINER fundon cfddream@gmail.com

RUN go get github.com/dockerboard/dockerboard
ENV GOPATH /go/src/github.com/dockerboard/dockerboard/Godeps/_workspace:/go
WORKDIR /go/src/github.com/dockerboard/dockerboard/

RUN go build dockerboard.go
RUN git clone --depth 1 https://github.com/dockerboard/bluewhale.git /bluewhale

EXPOSE 8001

ENTRYPOINT ["./dockerboard", "server"]
