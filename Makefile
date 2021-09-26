APP_NAME=golang-fiber-base-project
APP_VERSION=1.0.0
DOCKER_LOCAL=registry.gitlab.com/xdorro/registry
BUILD_DIR=./build
MAIN_DIR=./cmd/server

config:
	cp .env.example .env

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

docker.build:
	docker build  -f $(BUILD_DIR)/Dockerfile -t $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) .

docker.push:
	docker push $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION)

docker.run:
	docker run --name $(APP_NAME) -d -p 8000:8000 $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) -e SERVER.PORT=8000

docker.deploy: docker.build docker.run

docker.dev: docker.build docker.push

swag.init:
	swag init \
		-d $(MAIN_DIR) \
		-g main.go \
		-o ./api \
		--parseDependency \
		--parseInternal \
		--parseDepth 3

go.get:
	cd $(MAIN_DIR) && go get -u

go.gen:
	go generate ./...