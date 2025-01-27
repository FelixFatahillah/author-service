run:
	go run cmd/app/main.go


gen:
	mockgen -source=internal/service/interfaces.go -destination=internal/service/mocks/mock.go
	mockgen -source=internal/repository/interfaces.go -destination=internal/repository/mocks/mock.go

test:
	go test --short -coverprofile=tmp/cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=tmp/cover.out | grep "total"

generate-proto:
	protoc --proto_path=pkg/proto --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative pkg/proto/*.proto


run-client:
	go run cmd/client/main.go

run-server:
	go run cmd/server/main.go
