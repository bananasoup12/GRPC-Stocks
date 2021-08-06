# GRPC-Stocks
An API for creating, updating, and reading stock prices.

To recreate proto files, run 'protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     routeguide/route_guide.proto' from root dir.

To run, 
In one terminal or thread: go run server/server.go
in a second terminal: go run client/client.go