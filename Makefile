.PHONY: all build test bench bench_sync bench_async

all: build


build:
	go get github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate server -A service \
	 --exclude-main \
	 -t examples/service \
	 -f examples/service/swagger.yml
	go get ./...
	go build .

test:
	go test .
	go test -race -bench=BenchmarkDbAsync .


bench: bench_sync bench_async

bench_sync:
	go test -benchmem -bench=BenchmarkDbSync .

bench_async:
	go test -benchmem -bench=BenchmarkDbAsync .
