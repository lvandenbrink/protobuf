FROM golang:latest

COPY ./ /gopath/src/github.com/opsee/protobuf

WORKDIR $GOPATH/src/github.com/opsee/protobuf

RUN \
  go install ./opseeproto && \
  go install ./plugin/... && \
  go install ./protoc-gen-gogoopsee
