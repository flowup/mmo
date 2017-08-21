FROM golang:1.8

ADD intrfcs.go /go/src/intrfc.go
ADD mmo-plugin.yaml /plugin/mmo-plugin.yaml

RUN go build -o /hooks/intrfc /go/src/intrfc.go