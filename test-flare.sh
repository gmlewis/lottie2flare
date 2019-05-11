#!/bin/bash -ex

# Run once:
# jq -S . < examples/resizing-house/resizing-house.flr > examples/resizing-house/resizing-house.flr
# jq -S . < examples/simple-loader/simple-loader.flr > examples/simple-loader/simple-loader.flr

go run cmd/dump-flare/main.go \
  examples/resizing-house/resizing-house.flr \
  examples/simple-loader/simple-loader.flr

jq -S . < examples/resizing-house/resizing-house.flr.out.json > examples/resizing-house/data.out.json
diff -q examples/resizing-house/resizing-house.flr examples/resizing-house/data.out.json

jq -S . < examples/simple-loader/simple-loader.flr.out.json > examples/simple-loader/data.out.json
diff -q examples/simple-loader/simple-loader.flr examples/simple-loader/data.out.json
