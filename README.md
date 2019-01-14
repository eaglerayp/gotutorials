# Go-tutorial

* hw* are the folder for voluntary homework, can write the project under hw*/ and send merge request for peer review.
* tutorial*.md/pdf is the materials. using marp to generate pdf from md.

## Lesson 1

* Env setup
* DEV env setup (VSCode + linter + tools)
* go test
* golang common mistakes
* library exported concept
* common libraries
* gin

## Lesson 2

1. golang coding convention
2. golang package template
3. go example on go routines, channels
4. go profiling (including tracing, go test benchmark)
5. go mod

## HW

### HW1

* create a gin api server with middleware.
* extension: using any mongodb driver to create a CRUD api

### HW2

* split package from simple main.go, standardlize golang project layout
* try `go mod`
* try `context` timeout and using in a gin timeout api.
  * create a handler sleep 30 seconds, but return timeout error response in 15 seconds.