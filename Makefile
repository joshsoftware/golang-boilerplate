test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

bench:
	go test ./... -v -bench=. -benchmem