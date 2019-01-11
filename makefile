all: clean deps build test
buildProto:
	protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
build: buildProto
	go build -v -o target/grpc-test
buildDocker: buildProto
	GOOS=linux GARCH=amd64 go build -v -o target/grpc-test
clean:
	rm routeguide/route_guide.pb.go || true
	rm -rf target
deps:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/protoc-gen-go
	# needed for alpine compilation
	go get golang.org/x/sys/unix
docker: clean deps buildDocker test
	docker build -t shigglewitz/grpc:latest .
test:
	go test -v
