# User activity logger 

## Installation 
> go install 


## Running 
For REST Server
> go run cmd/server/main.go -type rest -port 8080

For GRPC Server
> go run cmd/server/main.go -type grpc -port 8081

For GRPC Client
> go run cmd/client/main.go 

Please use **postman** or **grpc cli** for testing server