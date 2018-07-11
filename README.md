# A potential Go 1.11-beta1 build regression

I found a build regression when building a large monorepo when trying out Go 1.11-beta1.

This project is the simplest reproducer I came up with.

## Go 1.10.3
```
$ go version
go version go1.10.3 darwin/amd64

$ go run main.go
OK
```

## Go 1.11-beta1

```
$ go version
go version go1.11beta1 darwin/amd64

$ go run main.go
# command-line-arguments
./main.go:11:6: undefined: addons.CustomJS
```

Looks like Go 1.11-beta1 ignores the [addons/custom_js.go](./addons/custom_js.go) file for some reason.

Is `_js.go` suffix somehow magical now because of the js/wasm support?

Is this intended? I don't think I saw this in the [Go 1.11 DRAFT release notes](https://tip.golang.org/doc/go1.11).
