FROM golang:1.8

ADD godeler /go/src/github.com/flowup/godeler
RUN go get github.com/flowup/godeler/... && go install github.com/flowup/godeler

ADD hooks/* /hooks/
ADD mmo-plugin.yaml /plugin/mmo-plugin.yaml

WORKDIR /source

CMD ["make"]