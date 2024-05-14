gen_protoc:
	protoc -I api api/user/user.proto --go_out=./pkg/gen --go_opt=paths=source_relative --go-grpc_out=./pkg/gen/ --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional

postgres:
	docker run --name user_service -e POSTGRES_PASSWORD=password  -p 5432:5432 --rm -d postgres

create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migrations -seq init

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations down

migrate-drop:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations drop