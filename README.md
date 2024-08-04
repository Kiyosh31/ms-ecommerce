# Overview

e-commerce microservices made in `go/golang`, `docker`, `kubernetes`, `gRPC`, `RabbitMQ`, `MongoDB`, `MongoDB Atlas`.

### Common package

You can find here the [common module](https://github.com/Kiyosh31/ms-ecommerce-common) used to share code and utils in this project

# Usage

### Pre requisites

Have installed:

- Go 1.22.5
- Docker desktop
- Kubernetes
- [gRPC and protobuf compiler](https://grpc.io/docs/languages/go/quickstart/) (just in case you want to develop new features)

## Instructions

1. clone the project
   ```console
   git@github.com:Kiyosh31/ms-ecommerce.git
   ```
2. Run docker and kubernetes

3. Create your own `secrets.yml` with all the required info

4. Open terminal in root project and run
   ```console
   make dev
   ```

## Links

- `RabbitMQ Management`: http://localhost:15672/

# Commands (Makefile)

- `make dev` It runs the project in developer mode
- `make dependency` It updates the common package for all microservices
- `make user/inventory/order/payment/product` compile and generates protobuf for specified service
- `make product` compile and generates protobuf for product-service
- `make clean` deletes all unused docker images generated by the project

# Documentation

Yo can find the `Postman` collection, design docs and some other documents on this repo's wiki once i finished the development
