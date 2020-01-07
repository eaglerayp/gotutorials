# Go-tutorial

* hw* are the folder for voluntary homework, can write the project under hw*/ and send merge request for peer review.

## Lesson 1

* Env setup
* DEV env setup (VSCode + linter + tools)
* go test
* go mod
* tips

## Lesson 2

* golang coding convention
* golang package template
* go example on go routines, channels
* library exported concept
* common libraries
* go context

## Lesson 3

* go unit test
* go benchmark
* go debug
* go profiling

## Lesson 4

* go error
* gin
* log
* kafka

## Lesson 5

* go grpc

## HW

### HW1

* create a gin api server with middleware.
* extension: using any mongodb driver to create a CRUD api

### HW2

* split package from simple main.go, standardlize golang project layout
* try `go mod`
* try `context` timeout and using in a gin timeout api.
  * create a handler sleep 30 seconds, but return timeout error response in 15 seconds.