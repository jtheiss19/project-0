#!/bin/bash

docker container prune -f

go build main.go init.go
docker build -t client-server -f ./Dockerfile .

cd Client/
go build Client.go
cd ..

cd Server/

cd LoadBalancer/
go build LoadBalancer.go
docker build -t load-balancer -f ./Dockerfile .
cd ..

cd ReverseProxy
go build ReverseProxy.go
docker build -t reverse-proxy -f ./Dockerfile .
cd ..

cd LoggingNode
go build LoggingNode.go
docker build -t logging-node -f ./Dockerfile .
cd ..

cd ..



