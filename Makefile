install_lint:
	go get -u golang.org/x/lint/golint

lint:
	golint ./...

test:
	go test -v ./...

benchmark:
	go test -v -bench=. -run=- ./...
