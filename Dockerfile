FROM alpine:latest

WORKDIR "/etc/grpc"

COPY target/grpc-test .

CMD ./grpc-test
