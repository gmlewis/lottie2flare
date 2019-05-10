#!/bin/bash -ex

# Run once:
# jq -S . < examples/adrock/data.json > examples/adrock/out.json
# jq -S . < examples/bodymovin/data.json > examples/bodymovin/out.json
# jq -S . < examples/gatin/data.json > examples/gatin/out.json
# jq -S . < examples/happy2016/data.json > examples/happy2016/out.json
# jq -S . < examples/navidad/data.json > examples/navidad/out.json

go run cmd/dump-lottie/main.go \
  examples/adrock/data.json \
  examples/bodymovin/data.json \
  examples/gatin/data.json \
  examples/happy2016/data.json \
  examples/navidad/data.json

jq -S . < examples/adrock/data.json.out.json > examples/adrock/data.out.json
diff -q examples/adrock/out.json examples/adrock/data.out.json
jq -S . < examples/bodymovin/data.json.out.json > examples/bodymovin/data.out.json
diff -q examples/bodymovin/out.json examples/bodymovin/data.out.json
jq -S . < examples/gatin/data.json.out.json > examples/gatin/data.out.json
diff -q examples/gatin/out.json examples/gatin/data.out.json
jq -S . < examples/happy2016/data.json.out.json > examples/happy2016/data.out.json
diff -q examples/happy2016/out.json examples/happy2016/data.out.json
jq -S . < examples/navidad/data.json.out.json > examples/navidad/data.out.json
diff -q examples/navidad/out.json examples/navidad/data.out.json
