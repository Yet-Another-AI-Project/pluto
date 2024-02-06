install:
	cd cmd/pluto-server && \
	go install -ldflags="-X 'main.VERSION=$(VERSION)'"

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t kiwihub.azurecr.io/pluto:$(VERSION) .

docker-build-staging:
	docker build --build-arg VERSION=staging -t kiwihub.azurecr.io/pluto:staging .

docker-push:
	docker push kiwihub.azurecr.io/pluto:$(VERSION)

docker-push-staging:
	docker push kiwihub.azurecr.io/pluto:staging

docker-clean:
	docker rmi kiwihub.azurecr.io/pluto:$(VERSION) || true
	docker rm -v $(shell docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null || true
	docker rmi $(shell docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null || true

docker-clean-staging:
	docker rmi kiwihub.azurecr.io/pluto:staging || true
	docker rm -v $(shell docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null || true
	docker rmi $(shell docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null || true

run: install
	pluto-server --config.file dev-config.yaml

server-binary-build:
	mkdir -p bin
	go build -ldflags="-X 'main.VERSION=$(VERSION)'" -o bin/pluto-server cmd/pluto-server/main.go

migrate-binary-build:
	mkdir -p bin
	go build -o bin/pluto-migrate cmd/pluto-migrate/main.go

unit-test:
	go test -v ./...

test: unit-test

ci-build-production: test docker-build docker-push docker-clean

ci-build-staging: test docker-build-staging docker-push-staging docker-clean-staging

generate-api-swagger:
	swag init -g cmd/pluto-server/main.go --output ./swagger
