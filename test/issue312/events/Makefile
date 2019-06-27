regenerate:
	go install github.com/FJSDS/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/FJSDS/protobuf/protobuf/ \
	--gogo_out=. events.proto 
