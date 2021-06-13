FROM golang:1.16.3-alpine
ENV GO111MODULE="on"
ENV CGO_ENABLED=0
WORKDIR /go-simple-mvc
COPY . /go-simple-mvc
RUN go build
