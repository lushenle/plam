# Image URL to use all building/pushing image targets
IMG ?= lushenle/plam:latest
DB_URL = postgresql://root:mypass@localhost:5432/plam?sslmode=disable

.PHONY: fmt
fmt:
	gofumpt -l -w .

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v -count=1 -cover --short ./...

.PHONY: test-coverage
test-coverage:
	go test -covermode=count -coverpkg=./... -coverprofile=coverage.out -v ./...
	go tool cover -html coverage.out -o coverage.html

.PHONY: server
server: fmt vet
	go run main.go

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: docker-build
docker-build:
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: docker
docker: docker-build docker-push

.PHONY: build
build:
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o plam main.go

.PHONY: mock
mock:
	mockgen -package mockdb -destination pkg/db/mock/store.go github.com/lushenle/plam/pkg/db Store

.PHONY: swagger
swagger: fmt
	swag fmt
	swag init --parseDependency --parseDepth 1 main.go --dir pkg/api --instanceName v1
