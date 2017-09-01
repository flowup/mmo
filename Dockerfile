FROM golang:1.9

MAINTAINER FlowUp <hello@flowup.cz>

# get dependency manager first
RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir /app
ADD . /go/src/github.com/flowup/mmo/
WORKDIR /go/src/github.com/flowup/mmo/

# fetch dependencies
RUN dep ensure

# build the binary
RUN go build -o mmo .

CMD ['/app/mmo']