docker run \
  -d \
  -v `pwd`/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg \
  -p 9000:80 \
  -p 9001:8001 \
  haproxy:1.5.19

docker run \
  -d \
  -e FUNCTION=server \
  -e GO_PORT=8080 \
  -p 8080:8080 \
  shigglewitz/grpc:latest

docker run \
  -d \
  -e FUNCTION=server \
  -e GO_PORT=8081 \
  -p 8081:8081 \
  shigglewitz/grpc:latest

docker run \
  -d \
  -e SERVER_ADDR=host.docker.internal:9001 \
  -e CONN_POLICY=refresh \
  -e GO_PORT=9006 \
  -p 9006:9006 \
  shigglewitz/grpc:latest

docker run \
  -d \
  -e SERVER_ADDR=host.docker.internal:9001 \
  -e CONN_POLICY=reuse \
  -e GO_PORT=9007 \
  -p 9007:9007 \
  shigglewitz/grpc:latest
