dev:
	skaffold dev

dependencies:
	cd gateway-api && go get -u
	cd user-service && go get -u

# Proto
USER_PROTO_PATH=./user-service/proto/user-service.proto
USER_PROTO_OUT_DIR=./user-service/proto/
GATEWAY_API_OUT_DIR=./gateway-api/generated/

user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(USER_PROTO_PATH)
	@cp $(USER_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)user-service