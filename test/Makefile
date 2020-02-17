generate:
	protoc --gogo_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --proto_path=$(GOPATH)/src:. testdata/*.proto

.PHONY:
	generate