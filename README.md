# Honey Badger

[![Go Report Card](https://goreportcard.com/badge/github.com/meeron/honey-badger)](https://goreportcard.com/report/github.com/meeron/honey-badger)
[![Go](https://github.com/meeron/honey-badger/actions/workflows/go-build.yml/badge.svg)](https://github.com/meeron/honey-badger/actions/workflows/go-build.yml)

Honey Badger is simple and fast cache server with persistent storage build on top of [BadgerDB](https://github.com/dgraph-io/badger). It uses [gRPC](https://grpc.io/) as transport protocol.

## Getting Started
### Build server
To build Honey Badger server you need [Go 1.21](https://go.dev/dl/) or above.

Windows users may need to install GNU Make. The best way is to use some package manager like [scoop](https://scoop.sh/#/apps?q=make)

To start, clone the repo

```sh
git clone git@github.com:meeron/honey-badger.git
```

Go to `honey-badger` directory and install packages

```sh
$ cd honey-badger
$ go install
```

Then run the following command
```sh
$ make build
```

This will produce server binary. Run it with default configuration
```sh
$ ./bin/hb
```

or on Windows
```
bin\hb.exe
```

### Docker

Run Docker image using
```sh
docker run --name honey-badger -p 18950:18950 -d meeron/honey-badger:latest
```

### Client
To connect to server you need gRPC client. Using [honey_badger.proto](https://github.com/meeron/honey-badger/blob/master/pb/honey_badger.proto) file you can generate one with your favorite [language](https://grpc.io/docs/languages/).

Check [server_test.go](https://github.com/meeron/honey-badger/blob/master/server/server_test.go) for examples in Go language.

To make call you can also use [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) command line tool
```sh
$ grpc_cli call localhost:18950 hb.Sys.Ping ""
connecting to localhost:18950
code: "pong"
Rpc succeeded with OK status
```

The command line tool also offers function to list avilable services on server
```sh
$ grpc_cli ls localhost:18950
hb.Data
hb.Db
hb.Sys
```
Then all methods in service
```sh
$ grpc_cli ls localhost:18950 hb.Data
Set
Get
GetByPrefix
Delete
DeleteByPrefix
SetBatch
```
Eventually you can print details for each previous commands.
```sh
$ grpc_cli ls localhost:18950 hb.Data.Set -l
rpc Set(hb.SetRequest) returns (hb.Result) {}
```

### Benchmark (in memory)
To test server performance on your system run
```sh
$ ./bin/hb -bench localhost:18950
os: darwin/arm64
cpus: 8

payload size: 256 bytes
num goroutines: 20
Set_30000: 293.691166ms
Set_50000: 509.71825ms
Set_100000: 1.026570458s

payload size: 256 bytes
num goroutines: 20
Get_30000: 317.108166ms
Get_50000: 489.316917ms
Get_100000: 982.176458ms

payload size: 256 bytes
num goroutines: 1
SetBatch_100000: 175.145125ms
SetBatch_300000: 596.115083ms
SetBatch_500000: 1.007408541s
```

The result `Set_30000: 293.691166ms` says that in 293ms 30k items has been sent to server
using 20 concurrent tasks.

## Hardware requirements
Honey Badger server should run on anything. CPU and RAM depends on your needs, but absolute minium is SSD disk (if persistance storage will be in use). Use benchmark command to check how Honey Badger is working on your instance.

## System requirements
### Linux and Mac
Honey Badger should build run on any Linux distro. [BadgerDB recommends](https://dgraph.io/docs/badger/faq/#are-there-any-linux-specific-settings-that-i-should-use) `max file descriptors` set to a high number depending upon the expected size of your data.

### Windows
It should build and run just fine.
