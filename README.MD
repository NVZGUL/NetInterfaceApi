# Network interface API and CLI
This programm provides a simple service that shows all the network interfaces available on the server

## Avaliable endpoints

```
GET services/version - get an API version
GET services/v1/interfaces - get list of avaliable network interfaces on server
GET services/v1/interfaces/%interface_name% - get detail information about interaface
```

## Running the program
To runnig the program use next command

```
docker-compose up --build -d 
```
and then
```
docker-compose run cli 
```
You will see a simple cli that provides access to API endpoints. CLI consist the next commands
```
--version - show the API version addres to services/version endpoint
help or h - show avaliable commands that avaliable in cli
list --server 127.0.0.1 --port 8080 - show the system and network interface list addres to services/v1/interfaces endpoint
show interface_name --server 127.0.0.1 --port 8080 - show detail information about interface in system address to services/v1/interfaces/%interface_name% endpoint
```