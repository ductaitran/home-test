FROM golang:1.18-alpine AS build

WORKDIR /go/build

COPY . ./

RUN apk add --no-cache git bash curl jq docker \
    && export PATH=/usr/local/go/bin:$PATH \
    && export GOPATH=/go \
    && export GOBIN=$GOPATH/bin \
    && go build -o $GOBIN/email-sender ./cmd

CMD ["bash"]