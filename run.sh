#!/bin/bash -ex
go run cmd/dump-lottie/main.go \
  examples/adrock/data.json \
  examples/bodymovin/data.json \
  examples/gatin/data.json \
  examples/happy2016/data.json \
  examples/navidad/data.json
