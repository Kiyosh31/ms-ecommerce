dev:
	skaffold dev

dependency:
	chmod +x update_common_dependency.sh
	./update_common_dependency.sh

# Proto
USER_PROTO_PATH=./user-service/proto/user-service.proto
USER_PROTO_OUT_DIR=./user-service/proto/
GATEWAY_API_OUT_DIR=./gateway-api/generated

user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(USER_PROTO_PATH)
	@rm -rf $(GATEWAY_API_OUT_DIR)/user-service
	@mkdir $(GATEWAY_API_OUT_DIR)/user-service
	@cp $(USER_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)/user-service



PRODUCT_PROTO_PATH=./product-service/proto/product-service.proto
PRODUCT_PROTO_OUT_DIR=./product-service/proto/

product:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PRODUCT_PROTO_PATH)
	@rm -rf $(GATEWAY_API_OUT_DIR)/product-service
	@mkdir $(GATEWAY_API_OUT_DIR)/product-service
	@cp $(PRODUCT_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)/product-service



INVENTORY_PROTO_PATH=./inventory-service/proto/inventory-service.proto
INVENTORY_PROTO_OUT_DIR=./inventory-service/proto/

inventory:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(INVENTORY_PROTO_PATH)
	@rm -rf $(GATEWAY_API_OUT_DIR)/inventory-service
	@mkdir $(GATEWAY_API_OUT_DIR)/inventory-service
	@cp $(INVENTORY_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)/inventory-service



ORDER_PROTO_PATH=./order-service/proto/order-service.proto
ORDER_PROTO_OUT_DIR=./order-service/proto/

order:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(ORDER_PROTO_PATH)
	@rm -rf $(GATEWAY_API_OUT_DIR)/order-service
	@mkdir $(GATEWAY_API_OUT_DIR)/order-service
	@cp $(ORDER_PROTO_OUT_DIR)*.pb.go $(GATEWAY_API_OUT_DIR)/order-service


# Clean docker
IMAGE_NAMES := gateway-api user-service product-service inventory-service order-service rabbitmq

clean:
	@for image in $(IMAGE_NAMES); do \
		echo "\nDeleting all Docker images with name $$image..."; \
		docker images | grep "$$image" | awk '{print $$3}' | xargs -r docker rmi -f || true; \
	done
	@echo "Docker images deleted."