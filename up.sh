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
