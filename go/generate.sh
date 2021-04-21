#!/bin/bash

set -e

protoc -I ../protobuf --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. \
  ../protobuf/types/*.proto \
  ../protobuf/types/time/*.proto \
  ../protobuf/info/*.proto \
  ../protobuf/traits/*.proto

go build ./...
go test ./...

