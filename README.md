# A potential Go 1.11-beta1 build regression

Building this project:

## Go 1.11-beta1

```
$ go version
go version go1.11beta1 darwin/amd64

$ go run main.go
# command-line-arguments
./main.go:11:6: undefined: addons.CustomJS
```

Looks like Go 1.11-beta1 ignores the [addons/custom_js.go](./addons/custom_js.go) file for some reason.


## Go 1.10.3
```
$ go version
go version go1.10.3 darwin/amd64

$ go run main.go
OK
```
