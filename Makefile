build:
	go build .

run:
	./cartesian-api

run-main:
	go run main.go

build-run:
	make build && make run

tests:
	go test  ./...

tests-coverage:
	go test -cover ./...