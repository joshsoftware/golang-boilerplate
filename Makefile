DB_HOST=localhost
DB_NAME=boilerplate_development
DB_USER=postgres
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/golang-boilerplate"
DB_PORT=5432

copy-config:
	cp application.yml.sample application.yml

build-deps:
	dep ensure -v

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

build: setup build-deps compile fmt vet lint

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint

db.setup: db.create db.migrate

db.create:
	createdb -h $(DB_HOST) -U $(DB_USER) -O $(DB_USER) $(DB_NAME)

db.drop:
	dropdb $(DB_NAME)

db.create_extension:
	psql -h $(DB_HOST) -d $(DB_NAME) -U $(DB_USER) -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp"';

db.migrate: db.create_extension
	$(APP_EXECUTABLE) migrate

db.reset: db.drop db.create db.migrate

test:
	go test -v ./...

test-coverage:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html
