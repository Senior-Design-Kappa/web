FROM golang:1.7.3-alpine

RUN apk add --update git gcc musl-dev

RUN export GOPATH=$HOME
RUN export SERVER_ADDRESS="159.203.88.91"

ADD . /go/src/github.com/Senior-Design-Kappa/web

RUN git clone https://github.com/Senior-Design-Kappa/web-client.git
RUN mv ./web-client /go/src/github.com/Senior-Design-Kappa/
RUN go install ./src/github.com/Senior-Design-Kappa/web

ENTRYPOINT /go/bin/web

EXPOSE 8000
