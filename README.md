1) make docker

2) ./up.sh

3) while true; do curl localhost:9006; done
the client running on this port will create a new grpc connection for every request, so you will see haproxy doing round robin load balancing

4) while true; do curl localhost:9007; done
the client running on this port will keep a single static connection, so there will be no round robin load balancing

5) ./down.sh
