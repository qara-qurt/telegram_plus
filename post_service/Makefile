mongo:
	docker run --rm --name mongodb -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=qaraqurt -d mongo:latest

gen_protoc:
	protoc -I api api/post/post.proto --go_out=./pkg/gen --go_opt=paths=source_relative --go-grpc_out=./pkg/gen/ --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional
