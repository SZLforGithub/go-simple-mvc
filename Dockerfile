FROM golang:1.16.3-alpine
ENV GO111MODULE="on"
ENV CGO_ENABLED=0
WORKDIR /go/src
COPY ./src/go.mod /go/src/
RUN go mod download
COPY ./src /go/src/
