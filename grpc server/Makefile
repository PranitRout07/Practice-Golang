build:
	protoc --go_out=./chat --go_opt=paths=source_relative --go-grpc_out=./chat --go-grpc_opt=paths=source_relative .\hello.proto

operations:
	protoc --go_out=./operations --go_opt=paths=source_relative --go-grpc_out=./operations --go-grpc_opt=paths=source_relative .\operations.proto