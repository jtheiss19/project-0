#!/bin/bash

go build main.go init.go
cd Client/
go build Client.go
cd ..
cd Server/
cd LoadBalancer/
go build LoadBalancer.go
cd ..
cd ReverseProxy
go build ReverseProxy.go
cd ..
cd ..
