# project-0
## Objective

Create a CLI which can create and manage many profiles. These profiles will contain health information and will help draw conclusions about the wellbeing of the profile owner's health. It will finally propose solutions to improving the overall health of the profile owner.

## Road Map

- [x] CLI can handle profiles
- [x] CLI can read profiles from file
- [x] CLI can save profiles to file
- [x] CLI can delete profiles
- [x] CLI can store Information in profiles
- [x] User can create profiles
- [x] User can delete profiles
- [x] User can input variables to be saved in specific profiles
- [x] User can overide information put into a profile in the past
- [x] User can request information stored in specific profiles
- [x] User can request the CLI program generate health information
- [x] User can request mass information from profiles
- [x] User can request CLI to draw health based conclusions
- [x] User can request CLI to give solutions to improve health

## How To Setup servers

Servers are connected to eachother through TCP connections. Users can use BuildAll.sh to build all main.go's located in the file structure and then run StartNetwork.sh to get the back components neaded for begin configuring a basic network.

### Clients

Clients are found in the client folder inside the main directory. They automatically try to make a connection to the client server through a port. If no connection can be formed they wait until a connection can be formed. 

### Reverse Proxy

The Reverse Proxy is found in the ReversaProxy folder in the Server folder inside the main directory. The reverse proxy when ran has no connections or redirections. Instead the user must input them manually at any point. They follow the form Host:Port. After that the Reverse Proxy will prompt the user for an identifier for the address they just entered. This identifier can be anything that the server it is redirecting too needs in its initial connection. Example for an http server identifier could be "GET" as the GET command is the first connection tcp connection. After entering an identifier the server can accept a new route. It can have many redirections. Redirections can be overridden by entering in the Host:Port name with a new identifier.

### Load Balancer

The Load Balancer is found in the LoadBalancer folder in the Server folder inside the main directory. The Load Balancer redircts traffic to many servers of the same kind. To add a new server once the load balancer is launched just use the command "add". It will then prompt the user to add the server by the form Host:Port. After entered it will now take any incomming connections and distribute them evenly to the servers in the backend. 

### Client Server

The Client server can be ran in the main directory by using ./main Host (PORT you want to host on). It allows access to all methods. This includes launching more servers if needed. Some methods are Show, Add, Edit, Remove, Del, etc.

### HTTP Server

The HTTP server can be ran in the main directory by using ./main Host -html. It serves the HTML located in the WebPages folder inside the Server folder inside the main directory. index.html is the landing page.