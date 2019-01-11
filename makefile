all: clean deps test build
build:
	go build -v -o target/grpc-test
clean:
	rm -rf target
deps:
	go get google.golang.org/grpc
test:
	go test -v
