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

# Clean docker
IMAGE_NAMES := gateway-api user-service

clean-images:
	@for image in $(IMAGE_NAMES); do \
		echo "Deleting all Docker images with name $$image...\n"; \
		docker images | grep "$$image" | awk '{print $$3}' | xargs -r docker rmi -f || true; \
	done
	@echo "Docker images deleted."