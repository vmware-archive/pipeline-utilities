FROM golang:1.9-alpine AS compilation

RUN apk update && apk add bash git unzip curl

ENV CGO_ENABLED=0
COPY . $GOPATH/src/github.com/calebwashburn/pipeline-utilities
RUN go get -a -t github.com/calebwashburn/pipeline-utilities/...
RUN go build github.com/calebwashburn/pipeline-utilities/cmd/pipeline-utilities

FROM alpine

RUN apk update && apk add bash unzip curl

COPY --from=compilation /go/bin/pipeline-utilities /usr/bin
