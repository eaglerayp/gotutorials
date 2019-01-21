<!-- $theme: gaia -->

# Golang Tutorial #3

* unit test
* benchmark
* go-grpc
* context

---

# unit test tool

* [assert](github.com/stretchr/testify/assert)
* [dockertest](gopkg.in/ory-am/dockertest)
* [gock](gopkg.in/h2non/gock)

---

# unit test flow control

```golang
// init_test.go
func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	var p *int
	retCode := 0
	p = &retCode
	BeforeTest()
	defer AfterTest(p)
	*p = m.Run()
}
```

---

# Assert

```golang
func TestGetMongoDBInfo(t *testing.T) {
	mongoConfig := getMongoDBInfo()
	assert.Equal(t, "testt", mongoConfig.Name)
}
```

```bash
--- FAIL: TestGetMongoDBInfo (0.00s)
    .../main_test.go:54:
        	Error Trace:	main_test.go:54
        	Error:      	Not equal:
        	            	expected: "testt"
        	            	actual  : "test"
        	Test:       	TestGetMongoDBInfo
FAIL
```

---

# HTTP mock

```golang
defer gock.Off() // Flush pending mocks after test execution
gock.InterceptClient(httpClient)
defer gock.RestoreClient(httpClient)
apDomain := "http://test.com"
path := "/test"
gock.New(apDomain).
    Get(path).
    Reply(200).
    JSON(map[string]string{
        "id": "123",
    })
```

---

# dockertest run mongo

```golang
var (
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource
)

dockerPool, err = dockertest.NewPool("")
dockerResource, err = dockerPool.Run("mongo", "3.4", nil)
dockerResource.GetPort("27017/tcp")
```

---

# dockertest teardown

```golang
func AfterTest(ret *int) {
	if e := recover(); e != nil {
		dockerPool.Purge(dockerResource)
		os.Exit(1)
	}
	dockerPool.Purge(dockerResource)
	os.Exit(*ret)
}
```

* sometimes teardown fail, please use `docker system prune -a`

---

# go benchmark #1

* `go test -benchmem -run=xxx`
* used when compared two or more syntax/function

```golang
func BenchmarkIfLt1(b *testing.B) {
	count := 0
	test := ""
	for n := 0; n < b.N; n++ {
		if len(test) < 1 {
			count++
		}
	}
	fmt.Println("lt1:", count)
}
```

---

# go benchmark #2


```golang
func BenchmarkIfEq0(b *testing.B) {
	count := 0
	test := ""
	for n := 0; n < b.N; n++ {
		if len(test) == 0 {
			count++
		}
	}
	fmt.Println("Eq0:", count)
}
```

---

# go gRPC implementaion

* [grpc intro](https://grpc.io/docs/guides/concepts.html)
* [go-grpc](https://github.com/grpc/grpc-go)
* `protoc --go_out=plugins=grpc:. *.proto`
* [go grpc example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)
* server and client struct implement interface
* `RegisgerXXXServiceServer` `NewXXXServiceClient`

---

# go context

* built-in library [context](https://golang.org/pkg/context/)

```golang
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}
```

---

# go context cancel

```golang
	ctx, cancel := context.WithCancel(context.Background())

    // after complete somethings...
    cancel()
```

---

# go context deadline/timeout

```golang
ctx, cancel := context.WithDeadline(context.Background(), d)

// Even though ctx will be expired, it is good practice to call its
// cancelation function in any case. Failure to do so may keep the
// context and its parent alive longer than necessary.
defer cancel()

select {
case <-time.After(1 * time.Second):
    fmt.Println("overslept")
case <-ctx.Done():
    fmt.Println(ctx.Err())
}
```