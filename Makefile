dev:
	go run .

build:
	go build -o ./bin/hb .

run: build
	./bin/hb

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/honey_badger.proto

bench:
	go build -o ./bin/hb_bench _bench/main.go
	./bin/hb_bench

test:
	go test ./... -v

