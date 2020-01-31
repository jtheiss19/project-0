#!/bin/bash

#Start Client Servers on 8090 and 8091
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/main Host 8090
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/main Host 8091

#Start Load Balancer
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Server/LoadBalancer/LoadBalancer

#Start Reverse Proxy
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Server/ReverseProxy/ReverseProxy

#Start Logging Server on 8070
x-terminal-emulator -e /home/joseph/go/src/github.com/jtheiss19/project-0/Server/LoggingNode/LoggingNode

