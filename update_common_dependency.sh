#!/bin/bash

# Define the repository URL
REPO_URL="github.com/Kiyosh31/ms-ecommerce-common"

# Fetch the latest tag from the GitHub repository
printf "\nGetting the list of tags...\n"
LATEST_TAG=$(git ls-remote --tags https://$REPO_URL | awk -F/ '{print $NF}' | sort -V | tail -n1)

# If fetching the latest tag fails, exit the script
if [ -z "$LATEST_TAG" ]; then
  printf "Failed to fetch the latest tag for $REPO_URL"
  exit 1
fi

# Use the latest tag with go get
printf "\nUpdating for /gateway-api..."
cd ./gateway-api
go get $REPO_URL@$LATEST_TAG
go mod tidy

printf "\nUpdating for ./user-service..."
cd ../user-service
go get $REPO_URL@$LATEST_TAG
go mod tidy

printf "\nUpdating for ./product-service..."
cd ../product-service
go get $REPO_URL@$LATEST_TAG
go mod tidy

printf "\nUpdating for ./inventory-service..."
cd ../inventory-service
go get $REPO_URL@$LATEST_TAG
go mod tidy

printf "\nUpdating for ./order-service..."
cd ../order-service
go get $REPO_URL@$LATEST_TAG
go mod tidy
