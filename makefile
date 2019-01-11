all: clean deps test build
build:
	protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
	go build -v -o target/grpc-test
clean:
	rm -rf target
deps:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
test:
	go test -v
