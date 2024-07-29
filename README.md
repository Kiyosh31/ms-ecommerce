# Overview

e-commerce microservices made in `go/golang`, `docker`, `kubernetes`, `gRPC`, `RabbitMQ`, `MongoDB`.

### Common package

You can find here the [common module](https://github.com/Kiyosh31/ms-ecommerce-common) used to share code and utils in this project

# Architecture

![ms-ecommerce](./img/ms-ecommerce-arch.png "a title")

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

3. Open terminal in root project and run
   ```console
   make run
   ```

# Commands (Makefile)

- `make dev` It runs the project in developer mode
- `make user` compile and generates protobuf for user-service
- `make product` compile and generates protobuf for product-service
- `make clean-images` deletes all unused docker images generated by the project

# Documentation

Yo can find the `Postman` collection and some other documents on this repo's wiki once i finished the development
