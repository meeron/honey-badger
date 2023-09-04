dev:
	go run .

build:
	go build -o ./bin/hb .
	cp config.json ./bin/config.json

run: build
	./bin/hb

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/honey_badger.proto

bench: build
	./bin/hb -bench

test:
	go test ./... -v

