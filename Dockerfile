FROM golang:alpine AS builder

COPY ./src $GOPATH/src

RUN cd "$GOPATH/src/go-hpa" \
 && go install

EXPOSE 80

ENTRYPOINT ["/go/bin/go-hpa"]