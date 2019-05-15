# flare-bug-workaround

There is a [known issue](https://github.com/2d-inc/Flare-Flutter/issues/90)
in the `flare_flutter` plugin that causes it to crash when the JSON keys are
in sorted order (which is the default for [json.Marshal](https://golang.org/pkg/encoding/json/#Marshal)).

I therefore wrote `flare-bug-workaround` as a temporary fix.

Usage:

```
$ go get github.com/gmlewis/lottie2flare/...
$ flare-bug-workaround file.flr
```

will write the file `file.flr.fix.flr` with the keys reordered so that the
`flare_flutter` plugin will no longer crash.
