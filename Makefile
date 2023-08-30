dev:
	go run .

build:
	go build -o ./.out/hb .

run: build
	./.out/hb

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/honey_badger.proto

bench:
	go test -bench=Benchmark

test:
	go test ./... -v

