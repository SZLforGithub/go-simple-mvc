FROM golang:1.16.3-alpine
ENV GO111MODULE="on"
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN go build -o app
