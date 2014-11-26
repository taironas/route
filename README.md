route
=====

taironas/route is an URL router in Go allowing usage of regexp in URL paths.

## Getting Started

After installing [Go](http://golang.org/doc/install) and setting up your [environment](http://golang.org/doc/code.html), create a `.go` file named `main.go`.

~~~ go
package main

import (
  "github.com/taironas/route"
  "net/http"
  "fmt"
)

func main() {
  r := new(route.Router)
  r.HandleFunc("/users/[0-9]+", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to page of user 42!")
  })

  http.ListenAndServe(":8080", r)
}
~~~

Then get the route package:
~~~
> go get github.com/taironas/route
~~~

Then build your server:
~~~
> cd $GOPATH/<APP_PATH>/<APP_DIR>
> go build
~~~

Then run your server:
~~~
> go run main
~~~

You will now have a Go net/http webserver running on `localhost:8080`.

## Testing

~~~
> go test
~~~

## Documentation

~~~
godoc -http=:6060
~~~
