dev:
	go run .

build:
	cd src && go build -o ./bin/hb -ldflags "-X main.version=$(ver)" .
	cp ./src/config.json ./bin/config.json

run: build
	./bin/hb

proto:
	protoc --go_out=./src --go_opt=paths=source_relative --go-grpc_out=./src --go-grpc_opt=paths=source_relative pb/honey_badger.proto

bench: build
	./bin/hb -bench 127.0.0.1:18950

test:
	cd src && go test ./... -v

docker:
	docker build --build-arg ver=$(ver) -t meeron/honey-badger:$(ver) -t meeron/honey-badger:latest .

