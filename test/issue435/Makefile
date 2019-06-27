regenerate:
	go install github.com/FJSDS/protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
