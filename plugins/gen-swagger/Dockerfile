FROM golang:1.8

RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get update && apt-get install -y --no-install-recommends nodejs unzip && rm -rf /var/lib/apt/lists/*

RUN mkdir proto && cd proto && \
    wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip && \
    unzip protoc-3.3.0-linux-x86_64.zip && \
    chmod +x bin/protoc

RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/gogo/protobuf/...
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN git clone https://github.com/googleapis/googleapis.git /googleapis

RUN npm install -g swagger-merger

ADD hooks/* /hooks/
ADD templates /plugin/templates
ADD mmo-plugin.yaml /plugin/mmo-plugin.yaml

ENV PROTO_PATH /go/proto/include
ENV PATH $PATH:/go/proto/bin

WORKDIR /source

CMD ["make"]