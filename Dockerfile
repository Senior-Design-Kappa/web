FROM golang:1.7.3-alpine

RUN apk add --update git gcc musl-dev

RUN export GOPATH=$HOME
RUN export SERVER_ADDRESS="159.203.88.91"

RUN git clone https://github.com/Senior-Design-Kappa/web.git

ADD . /go/src/github.com/Senior-Design-Kappa/web

RUN go install ./src/github.com/Senior-Design-Kappa/web

ENTRYPOINT /go/bin/web

EXPOSE 8000
