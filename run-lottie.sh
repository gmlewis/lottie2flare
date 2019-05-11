#!/bin/bash -ex

# Run once:
# jq -S . < examples/adrock/data.json > examples/adrock/out.json
# jq -S . < examples/bodymovin/data.json > examples/bodymovin/out.json
# jq -S . < examples/gatin/data.json > examples/gatin/out.json
# jq -S . < examples/happy2016/data.json > examples/happy2016/out.json
# jq -S . < examples/navidad/data.json > examples/navidad/out.json
# jq -S . < examples/1055-world-locations/1055-world-locations.json > examples/1055-world-locations/out.json
# jq -S . < examples/113-muzli-beacon/113-muzli-beacon.json > examples/113-muzli-beacon/out.json
# jq -S . < examples/117-progress-bar/117-progress-bar.json > examples/117-progress-bar/out.json
# jq -S . < examples/1370-confetti/1370-confetti.json > examples/1370-confetti/out.json
# jq -S . < examples/1798-check-animation/1798-check-animation.json > examples/1798-check-animation/out.json
# jq -S . < examples/196-material-wave-loading/196-material-wave-loading.json > examples/196-material-wave-loading/out.json
# jq -S . < examples/201-simple-loader/201-simple-loader.json > examples/201-simple-loader/out.json
# jq -S . < examples/257-favorie/257-favorie.json > examples/257-favorie/out.json
# jq -S . < examples/27-loading/27-loading.json > examples/27-loading/out.json
# jq -S . < examples/29-motorcycle/29-motorcycle.json > examples/29-motorcycle/out.json
# jq -S . < examples/40-loading/40-loading.json > examples/40-loading/out.json
# jq -S . < examples/427-happy-birthday/427-happy-birthday.json > examples/427-happy-birthday/out.json
# jq -S . < examples/433-checked-done/433-checked-done.json > examples/433-checked-done/out.json
# jq -S . < examples/51-preloader/51-preloader.json > examples/51-preloader/out.json
# jq -S . < examples/72-favourite-app-icon/72-favourite-app-icon.json > examples/72-favourite-app-icon/out.json
# jq -S . < examples/77-im-thirsty/77-im-thirsty.json > examples/77-im-thirsty/out.json
# jq -S . < examples/782-check-mark-success/782-check-mark-success.json > examples/782-check-mark-success/out.json

go run cmd/dump-lottie/main.go \
  examples/adrock/data.json \
  examples/bodymovin/data.json \
  examples/gatin/data.json \
  examples/happy2016/data.json \
  examples/navidad/data.json \
  examples/1055-world-locations/1055-world-locations.json \
  examples/113-muzli-beacon/113-muzli-beacon.json \
  examples/117-progress-bar/117-progress-bar.json \
  examples/1370-confetti/1370-confetti.json \
  examples/1798-check-animation/1798-check-animation.json \
  examples/196-material-wave-loading/196-material-wave-loading.json \
  examples/201-simple-loader/201-simple-loader.json \
  examples/257-favorie/257-favorie.json \
  examples/27-loading/27-loading.json \
  examples/29-motorcycle/29-motorcycle.json \
  examples/40-loading/40-loading.json \
  examples/427-happy-birthday/427-happy-birthday.json \
  examples/433-checked-done/433-checked-done.json \
  examples/51-preloader/51-preloader.json \
  examples/72-favourite-app-icon/72-favourite-app-icon.json \
  examples/77-im-thirsty/77-im-thirsty.json \
  examples/782-check-mark-success/782-check-mark-success.json

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

jq -S . < examples/1055-world-locations/1055-world-locations.json.out.json > examples/1055-world-locations/data.out.json
diff -q examples/1055-world-locations/out.json examples/1055-world-locations/data.out.json

jq -S . < examples/113-muzli-beacon/113-muzli-beacon.json.out.json > examples/113-muzli-beacon/data.out.json
diff -q examples/113-muzli-beacon/out.json examples/113-muzli-beacon/data.out.json

jq -S . < examples/117-progress-bar/117-progress-bar.json.out.json > examples/117-progress-bar/data.out.json
diff -q examples/117-progress-bar/out.json examples/117-progress-bar/data.out.json

jq -S . < examples/1370-confetti/1370-confetti.json.out.json > examples/1370-confetti/data.out.json
diff -q examples/1370-confetti/out.json examples/1370-confetti/data.out.json

jq -S . < examples/1798-check-animation/1798-check-animation.json.out.json > examples/1798-check-animation/data.out.json
diff -q examples/1798-check-animation/out.json examples/1798-check-animation/data.out.json

jq -S . < examples/196-material-wave-loading/196-material-wave-loading.json.out.json > examples/196-material-wave-loading/data.out.json
diff -q examples/196-material-wave-loading/out.json examples/196-material-wave-loading/data.out.json

jq -S . < examples/201-simple-loader/201-simple-loader.json.out.json > examples/201-simple-loader/data.out.json
diff -q examples/201-simple-loader/out.json examples/201-simple-loader/data.out.json

jq -S . < examples/257-favorie/257-favorie.json.out.json > examples/257-favorie/data.out.json
diff -q examples/257-favorie/out.json examples/257-favorie/data.out.json

jq -S . < examples/27-loading/27-loading.json.out.json > examples/27-loading/data.out.json
diff -q examples/27-loading/out.json examples/27-loading/data.out.json

jq -S . < examples/29-motorcycle/29-motorcycle.json.out.json > examples/29-motorcycle/data.out.json
diff -q examples/29-motorcycle/out.json examples/29-motorcycle/data.out.json

jq -S . < examples/40-loading/40-loading.json.out.json > examples/40-loading/data.out.json
diff -q examples/40-loading/out.json examples/40-loading/data.out.json

jq -S . < examples/427-happy-birthday/427-happy-birthday.json.out.json > examples/427-happy-birthday/data.out.json
diff -q examples/427-happy-birthday/out.json examples/427-happy-birthday/data.out.json

jq -S . < examples/433-checked-done/433-checked-done.json.out.json > examples/433-checked-done/data.out.json
diff -q examples/433-checked-done/out.json examples/433-checked-done/data.out.json

jq -S . < examples/51-preloader/51-preloader.json.out.json > examples/51-preloader/data.out.json
diff -q examples/51-preloader/out.json examples/51-preloader/data.out.json

jq -S . < examples/72-favourite-app-icon/72-favourite-app-icon.json.out.json > examples/72-favourite-app-icon/data.out.json
diff -q examples/72-favourite-app-icon/out.json examples/72-favourite-app-icon/data.out.json

jq -S . < examples/77-im-thirsty/77-im-thirsty.json.out.json > examples/77-im-thirsty/data.out.json
diff -q examples/77-im-thirsty/out.json examples/77-im-thirsty/data.out.json

jq -S . < examples/782-check-mark-success/782-check-mark-success.json.out.json > examples/782-check-mark-success/data.out.json
diff -q examples/782-check-mark-success/out.json examples/782-check-mark-success/data.out.json
