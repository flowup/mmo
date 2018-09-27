FROM golang:1.11 as build

WORKDIR /{{ .Name }}
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/{{ .Name }} ./cmd/{{ .Name }}

FROM alpine:latest

WORKDIR /app

COPY --from=0 /go/bin/{{ .Name }} .

CMD ["/app/{{ .Name }}"]
