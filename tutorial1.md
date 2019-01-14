<!-- $theme: gaia -->

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

# Libraries

* context https://golang.org/pkg/context/
* log: https://github.com/sirupsen/logrus
* exported package concept
* gin github.com/gin-gonic/gin
* https://mholt.github.io/json-to-go/

---

# Tips

* fmt.Sprintf() https://golang.org/pkg/fmt/
* https://blog.golang.org/defer-panic-and-recover
* json marshal by struct & tag
* use `` declare json string which including "