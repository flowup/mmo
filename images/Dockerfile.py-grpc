FROM python:3.6

RUN apt-get update && apt-get install -y --no-install-recommends unzip git && rm -rf /var/lib/apt/lists/*

RUN mkdir /proto && cd /proto && \
    wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip && \
    unzip protoc-3.3.0-linux-x86_64.zip && \
    rm -rf /proto/protoc-3.3.0-linux-x86_64.zip && \
    chmod +x bin/protoc

RUN pip install grpcio-tools

RUN git clone https://github.com/googleapis/googleapis.git /googleapis

ENV PROTO_PATH /proto/include
ENV PATH $PATH:/proto/bin

WORKDIR /in

CMD ["make"]