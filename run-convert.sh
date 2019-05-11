#!/bin/bash -ex

go run cmd/lottie2flare/main.go \
  examples/201-simple-loader/201-simple-loader.json

jq -S . < examples/201-simple-loader/201-simple-loader.json.flr > examples/201-simple-loader/convert.json
diff -q examples/simple-loader/simple-loader.flr examples/201-simple-loader/convert.json
