# -- docker builds
IMAGE_SERVER ?= product-api-server:latest
build-server:
	docker build . --build-arg cmd=server -t $(IMAGE_SERVER)

# -- docker compose
start: build-server
	docker-compose up -d

stop:
	docker-compose stop

clean: stop
	docker system prune --all --volumes

# -- dev tools
lint:
	golangci-lint run

# -- backend seeders
install-seeders-tools:
	brew install curl
	brew install jq

upload-fees:
	@curl -X POST \
      http://localhost:5000/fees/massive-upload \
      -H 'authorization: $(ADMIN_TOKEN)' \
      -H 'content-type: multipart/form-data' \
      -F file=@./internal/api/testing/assets/fees.xlsx
