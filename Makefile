dev:
	skaffold dev

dependencies:
	cd gateway-api && go get -u github.com/Kiyosh31/ms-ecommerce-common && go mod tidy
	cd user-service && go get -u github.com/Kiyosh31/ms-ecommerce-common && go mod tidy

# Proto
USER_PROTO_PATH=./user-service/proto/user-service.proto
USER_PROTO_OUT_DIR=./user-service/proto/
GATEWAY_API_OUT_DIR=./gateway-api/generated/

user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(USER_PROTO_PATH)
	@cp $(USER_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)user-service