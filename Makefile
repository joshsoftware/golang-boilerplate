copy-config:
	cp application.yml.sample application.yml

test:
	go test -v ./...
