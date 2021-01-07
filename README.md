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


## Client commands 

> go run cmd/client/main.go create-user --email "test@gmail.com" --phone-no 12312321312 --username "asdasd" --server-addr 0.0.0.0:8080

> go run cmd/client/main.go add-activity --server-addr 0.0.0.0:8080 --email test@gmail.com --activity-day 2020-10-10 --activity-status active --activity-time-duration 7 --activity-type sleep

> go run cmd/client/main.go get-activities --server-addr 0.0.0.0:8080 --email test@gmail.com

> go run cmd/client/main.go update-activity --server-addr 0.0.0.0:8080 --email test@gmail.com --activity-day 2020-10-10 --activity-status done --activity-type sleep 

// Check update activities 
> go run cmd/client/main.go get-activities --server-addr 0.0.0.0:8080 --email test@gmail.com
