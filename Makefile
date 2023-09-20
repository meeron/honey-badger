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
	./bin/hb -bench 127.0.0.1:18950

test:
	go test ./... -v

docker:
	docker build -t meeron/honey-badger:0.1.0-alpha.1 .

