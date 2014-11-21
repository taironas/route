// taironas/route is an URL router in Go allowing usage of regexp in URL paths.
//
//  package main
//
//  import (
//    "github.com/taironas/route"
//    "net/http"
//    "fmt"
//  )
//
//  func main() {
//    r := new(route.Router)
//    r.HandleFunc("/users/[0-9]+", func(w http.ResponseWriter, req *http.Request) {
//      fmt.Fprintf(w, "Welcome to page of user 42!")
//    })
//
//    http.ListenAndServe(":8080", r)
//  }
package route
