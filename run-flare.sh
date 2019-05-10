#!/bin/bash -ex

# Run once:
jq -S . < examples/resizing-house/resizing-house.flr > examples/resizing-house/resizing-house.json

go run cmd/dump-flare/main.go \
  examples/resizing-house/resizing-house.flr

jq -S . < examples/resizing-house/resizing-house.flr.out.json > examples/resizing-house/data.out.json
diff -q examples/resizing-house/resizing-house.json examples/resizing-house/data.out.json
