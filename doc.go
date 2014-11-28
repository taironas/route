// Package route provides URL router allowing usage of regexp in URL paths.
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
//    r.HandleFunc("/users/?", usersHandler)
//    r.HandleFunc("/users/[0-9]+", userHandler)
//    r.HandleFunc("/users/[0-9]+/friends/[a-zA-Z]+", friendHandler)
//
//    http.ListenAndServe(":8080", r)
//  }
//
//  func usersHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to users handler!")
//  }
//
//  func userHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to user handler!")
//  }
//
//  func friendHandler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Welcome to friend handler!")
//  }
package route
