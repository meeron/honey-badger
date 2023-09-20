dev:
	go run .

build:
	go build -o ./bin/hb -ldflags "-X main.version=$(ver)" .
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
	docker build --build-arg ver=$(ver) -t meeron/honey-badger:$(ver) -t meeron/honey-badger:latest .

