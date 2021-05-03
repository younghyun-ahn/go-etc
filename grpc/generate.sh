protoc hello/hellopb/hello.proto --go_out=plugins=grpc:.

# go run hello/hello_server/server.go
# go run hello/hello_client/client.go