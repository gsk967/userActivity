
gen:
	protoc --proto_path=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --openapiv2_out=:swagger \
    ./proto/*

clean:
	rm -r pb/*

grpc:
	 go run cmd/server/main.go -type grpc -port 8080

rest:
	go run cmd/server/main.go -type rest -port 8081

client:
	go run cmd/server/client.go -addr 0.0.0.0:8080

test:
	go test -cover -race ./...