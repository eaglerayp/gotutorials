---
marp: true
---

<!-- theme: gaia -->

# Golang Tutorial #1

## Getting started

---

# Environment Setup

* env GOROOT/GOPATH/PATH set in ~/.bashrc
* `rm -rf /usr/local/go` if update big version

---

# Dev Env setup (VSCODE)

* plugin: go, install go anylasis tools
* https://github.com/eaglerayp/DevTools
* github.com/golangci/golangci-lint
* go build -i (make autocomplete work)
* go test package

---

# Tools

* go test cmd https://golang.org/cmd/go/#hdr-Testing_flags
* goDoc https://godoc.org/golang.org/x/tools/cmd/godoc

---

# go mod init

* `go mod init github.com/xxx/xxx` only create go.mod
* `go build ./...` generate items in go.mod & go.sum
* force version by set `go.mod`
* export GOPRIVATE=gitlab.com* for private code base
* rm go.sum first to avoid append

```mod
module gitlab.com/eaglerayp/...
go 1.13
require (
	cloud.google.com/go v0.34.0 // indirect
```

---

# go mod build/test

* `go build, go test` will automatically add new dependencies(updating go.mod and downloading the new dependencies).
* create vendor `go mod vendor`, build by vendor `go build/install -mod=vendor`
* it's easier to use vendor if there is private dependency.

---

# Tips

* fmt.Sprintf() https://golang.org/pkg/fmt/
* https://blog.golang.org/defer-panic-and-recover
* json marshal by struct & tag
* use `` declare json string which including "
