#!/bin/bash

#Start Client Servers
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/main Host 8090
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/main Host 8091

#Start Load Balancer on 8082
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Server/LoadBalancer/LoadBalancer

#Start Reverse Proxy on 8081
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Server/ReverseProxy/ReverseProxy

#Launch Three clients
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Client/Client