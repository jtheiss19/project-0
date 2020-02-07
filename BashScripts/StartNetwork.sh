#!/bin/bash

docker container prune -f

#Start Client Servers on 8090 and 8091
x-terminal-emulator -e docker container run --publish 8090:8080 --name ClientServer1 -rm -i client-server 
x-terminal-emulator -e docker container run --publish 8091:8080 --name ClientServer2 -rm -i client-server

#Start Load Balancer
x-terminal-emulator -e  docker container run --publish 8081:8080 --name Load1 -rm -i load-balancer

#Start Reverse Proxy
x-terminal-emulator -e  docker container run --publish 8082:8080 --name Proxy1 -rm -i reverse-proxy

#Start Logging Server on 8070
x-terminal-emulator -e  docker container run --publish 8070:8080 --name Log1 -rm -i logging-node

