dev:
	skaffold dev

dependencies:
	cd gateway-api && go get -u github.com/Kiyosh31/ms-ecommerce-common
	cd user-service && go get -u github.com/Kiyosh31/ms-ecommerce-common

user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    user-service/proto/user-service.proto