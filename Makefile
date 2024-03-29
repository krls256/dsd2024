gen-logging-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
	 --go-grpc_opt=require_unimplemented_servers=false \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative api/logging.proto

gen-messages-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
	 --go-grpc_opt=require_unimplemented_servers=false \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative api/messages.proto

run-logging:
	go run ./cmd/logging

run-messages:
	go run ./cmd/messages

run-facade:
	go run ./cmd/facade