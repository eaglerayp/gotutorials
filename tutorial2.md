<!-- $theme: gaia -->

# Golang Tutorial #2

## Hands on

---

# golang coding convention

* camelCase
* [official blog](https://github.com/golang/go/wiki/CodeReviewComments)
* package name 單數名詞
* package 內部 variable name 不要重複 prefix, e.g., `chubby.ChubbyFile`
* [error variable naming](https://github.com/golang/go/wiki/Errors): prefix with `err` or `Err`
* Named Result Parameters 個人喜好

---

# golang package template/example

* example: [awesome](https://github.com/avelino/awesome-go)
  * [server application](https://github.com/hashicorp/consul)
  * [cmd tools](https://github.com/drone/drone)
  * [library](https://github.com/gin-gonic/gin)
* [project layout](https://github.com/golang-standards/project-layout)

---

# go example on go routines, channels 用法

* https://blog.golang.org/go-concurrency-patterns-timing-out-and
* https://blog.golang.org/pipelines
* https://talks.golang.org/2012/concurrency.slide
* Channel Example:  https://gitlab.devops.maaii.com/general-backend/mongodao/blob/master/pool.go#L75
* Goroutine Example: https://gitlab.devops.maaii.com/general-backend/m800-application-plugin/blob/master/controller/gin.go#L446

---

# go mod

* `go mod init github.com/xxx/xxx`

```mod
module gitlab.com/eaglerayp/lotushouse

require (
	cloud.google.com/go v0.34.0 // indirect
	firebase.google.com/go v3.5.0+incompatible
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	go.opencensus.io v0.18.0 // indirect
	golang.org/x/oauth2 v0.0.0-20181203162652-d668ce993890 // indirect
	google.golang.org/api v0.0.0-20181217000635-41dc4b66e69d
	google.golang.org/grpc v1.17.0 // indirect
)
```

* build `go build ./...`
 * Standard commands like go build or go test will automatically add new dependencies as needed to satisfy imports (updating go.mod and downloading the new dependencies).
* create vendor `go mod vendor`

---

#  go profiling (including tracing, go test benchmark)

* https://github.com/davecheney/gophercon2018-performance-tuning-workshop
* `go tool pprof`
* [http pprof](https://golang.org/pkg/net/http/pprof/)
* `go tool pprof -http=":8011" http://localhost:10201/debug/pprof/profile?seconds=30`
* http pprof example: https://gitlab.devops.maaii.com/general-backend/m800-application-plugin/blob/master/prof.go

---

# Extend reading

* [Error handling and Go](https://blog.golang.org/error-handling-and-go)
* [Go Errors](https://dave.cheney.net/paste/gocon-spring-2016.pdf)
* [Visualize go routines](https://divan.github.io/posts/go_concurrency_visualize/)