# faasify - A framework for building a simple and blazingly fast FaaS server from scratch

## Compile the server with a security token

    ./scripts/compile.sh

## Run the server in development mode

Use the default (domain=localhost, url=https://localhost:3000)

    ./build/server 

Or specify the domain and url with args:

    ./build/server -domain localhost -url https://localhost:3000

## Call a function with curl and security token

    curl -H "Authorization: Bearer $(cat ./security/token)" https://localhost:3000/hello

## Display the demo page

[https://127.0.0.1:3000/index](https://127.0.0.1:3000/index)

Login with the username <code>faasify</code> and the security token in <code>security/token</code> as the password.

## Add functions

Create a new function named <code>count</code>:

    mkdir ./functions/count

Add a function named <code>HandlerFunc()</code> like the following:
    
    vim ./functions/count/handler.go

```go

package count

type Request struct{}
type Response struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}

var (
	count int
	mutex sync.Mutex
)

func HandlerFunc(cfg *config.Config) http.HandlerFunc {
	count = 0
	return func(w http.ResponseWriter, r *http.Request) {
		server.Process[Request, Response](w, r, func(req Request) (Response, error) {
			mutex.Lock()
			defer mutex.Unlock()
			count++
			return Response{Count: count}, nil
		})
	}
}
```

Add a handler test on steroids like this:

```go

package count_test

func TestCountSuccess(t *testing.T) {
	cfg := &config.Config{}
	fn := count.HandlerFunc(cfg)
	req := count.Request{}
	res, err := server.Validate[count.Request, count.Response](fn, "count", req, cfg)
	assert.That("no error is returned", t, err, nil)
	assert.That("count is 1", t, res.Count, 1)
}

```

The [function routing](/internal/http/server/router.go) gets updated on compile time, based on the contents of the functions directory.

## Run the server in production mode

Start server with TLS support:

	sudo ./build/server -domain your.domain -mode prod -url https://your.domain

