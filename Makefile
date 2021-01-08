
gen:
	protoc --proto_path=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --openapiv2_out=:swagger \
    ./proto/*

clean:
	rm -r pb/*

install:
	go install ./...
	go build -o ./build/clientd -ldflags "-s -w" ./cmd/client/main.go
	go build -o ./build/userActivity -ldflags "-s -w" ./cmd/server/main.go

grpc:
	./build/userActivity -type grpc -port 8080

rest:
	./build/userActivity -type rest -port 8081

client:
	./build/clientd -addr 0.0.0.0:8080

test:
	go test -cover -race ./...