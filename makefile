all: clean deps build test
build:
	protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
	go build -v -o target/grpc-test
clean:
	rm routeguide/route_guide.pb.go || true
	rm -rf target
deps:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
test:
	go test -v
