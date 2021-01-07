
gen:
	protoc --proto_path=./proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --openapiv2_out=:swagger \
    ./proto/*

clean:
	rm -r pb/*

server:
	 go run cmd/server/main.go

client:
	echo "Client...."