# faasify - A framework for building a simple and blazingly fast FaaS server from scratch

## Compile the server with a security token

    ./scripts/compile.sh

## Run the server in development mode

    sudo ./build/server -url http://localhost

## Call a function with curl and security token

    curl -H "Authorization: Bearer $(cat ./security/token)" http://localhost/status

## Call a function with faasify client and buildin token

    ./build/client -host http://localhost -name status

## Display the demo page

[http://127.0.0.1/index](http://127.0.0.1/index)

## Add functions

Create a new function named <code>YOUR_FUNCTION</code>:

    mkdir ./functions/YOUR_FUNCTION

Add a function named <code>HandlerFunc()</code> like the follow:
    
    vim ./functions/YOUR_FUNCTION/handler.go

```go

package YOUR_FUNCTION

import "net/http"

func HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}
```

The [function routing](/internal/http/server/router.go) gets updated on compile time, based on the contents of the functions directory.

## Run the server in production mode

Start server with TLS support:

	sudo ./build/server -domain your.domain -mode prod -url https://your.domain

