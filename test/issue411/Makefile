regenerate:
	go install github.com/FJSDS/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/FJSDS/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue411.proto
