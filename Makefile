install_lint:
	go install golang.org/x/lint/golint@latest

lint:
	golint ./...

test:
	go test -v ./...

benchmark:
	go test -v -bench=. -run=- ./...
