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

