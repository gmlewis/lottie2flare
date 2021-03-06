# Read/Write lottie and Flutter flare files in Go

This is an experimental library and tool used to convert
[lottie-web](https://github.com/airbnb/lottie-web) files
to [Flutter](https://flutter.dev) [flare](https://2dimensions.com/flare) files.

It is currently a work-in-progress.
Only an infinitesimal subset of simple animations are currently supported.

Note that there is a [known issue](https://github.com/2d-inc/Flare-Flutter/issues/90)
in the `flare_flutter` plugin that causes it to crash when the JSON keys are
in sorted order (which is the default for [json.Marshal](https://golang.org/pkg/encoding/json/#Marshal)).
I therefore wrote [flare-bug-workaround](cmd/flare-bug-workaround) as a temporary fix.

![201-simple-loader](examples/201-simple-loader/201-simple-loader.gif)

## Examples

Please see the examples:

* [adrock](examples/adrock)
* [bodymovin](examples/bodymovin)
* [gatin](examples/gatin)
* [happy2016](examples/happy2016)
* [navidad](examples/navidad)
* [1055-world-locations](examples/1055-world-locations)
* [113-muzli-beacon](examples/113-muzli-beacon)
* [117-progress-bar](examples/117-progress-bar)
* [1370-confetti](examples/1370-confetti)
* [1798-check-animation](examples/1798-check-animation)
* [196-material-wave-loading](examples/196-material-wave-loading)
* [201-simple-loader](examples/201-simple-loader)
* [257-favorie](examples/257-favorie)
* [27-loading](examples/27-loading)
* [29-motorcycle](examples/29-motorcycle)
* [40-loading](examples/40-loading)
* [427-happy-birthday](examples/427-happy-birthday)
* [433-checked-done](examples/433-checked-done)
* [51-preloader](examples/51-preloader)
* [72-favourite-app-icon](examples/72-favourite-app-icon)
* [77-im-thirsty](examples/77-im-thirsty)
* [782-check-mark-success](examples/782-check-mark-success)


## Documentation
[![GoDoc](https://godoc.org/github.com/gmlewis/lottie2flare/lottie?status.svg)](https://godoc.org/github.com/gmlewis/lottie2flare/lottie)

----------------------------------------------------------------------

Enjoy!

----------------------------------------------------------------------

# License

Copyright 2019 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
