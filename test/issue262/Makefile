regenerate:
	go install github.com/FJSDS/protobuf/protoc-min-version
	go install github.com/FJSDS/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/FJSDS/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/FJSDS/protobuf/types:. timefail.proto
