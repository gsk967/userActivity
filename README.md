# User activity logger 

## Installation 
> go install ./...


## Running 
Install 
> make install 

For REST Server
> ./build/userActivity  -type rest -port 8080

For GRPC Server
> ./build/userActivity -type grpc -port 8081

For GRPC Client
> ./build/clientd --server-addr 0.0.0.0:8080

Please use **postman** or **grpc cli** for testing server


## Client commands 

> ./build/clientd create-user --email "test@gmail.com" --phone-no 12312321312 --username "asdasd" --server-addr 0.0.0.0:8080

> ./build/clientd add-activity --server-addr 0.0.0.0:8080 --email test@gmail.com --activity-day 2020-10-10 --activity-status active --activity-time-duration 7 --activity-type sleep

> ./build/clientd get-activities --server-addr 0.0.0.0:8080 --email test@gmail.com

> ./build/clientd update-activity --server-addr 0.0.0.0:8080 --email test@gmail.com --activity-day 2020-10-10 --activity-status done --activity-type sleep 

// Check update activities 
> ./build/clientd get-activities --server-addr 0.0.0.0:8080 --email test@gmail.com
