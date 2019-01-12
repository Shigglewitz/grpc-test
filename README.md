1) make docker

2) ./up.sh

3) while true; do curl localhost:9006; done

3a) To see round robin in haproxy take effect, kill the docker container running the client (bound to port 9006) and restart it.  Then responses will start coming from port 8001

4) ./down.sh
