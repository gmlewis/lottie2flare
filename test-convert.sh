#!/bin/bash -ex

go run cmd/lottie2flare/main.go \
  examples/201-simple-loader/201-simple-loader.json \
  examples/27-loading/27-loading.json

jq -S . < examples/201-simple-loader/201-simple-loader.json.flr > examples/201-simple-loader/convert.json
diff -q examples/201-simple-loader/want.flr examples/201-simple-loader/convert.json || \
  diff examples/201-simple-loader/want.flr examples/201-simple-loader/convert.json

jq -S . < examples/27-loading/27-loading.json.flr > examples/27-loading/convert.json
diff -q examples/27-loading/want.flr examples/27-loading/convert.json || \
  diff examples/27-loading/want.flr examples/27-loading/convert.json
