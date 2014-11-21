package route

import (
	"fmt"
  "net/http"
	"net/http/httptest"
  "testing"
)

var endpoints = []string{"/test/handler/1", "/test/handler/2"}

func TestRouter(t *testing.T) {
  r := new(Router)
	r.HandleFunc(endpoints[0], testHandler1)
  r.HandleFunc(endpoints[1], testHandler2)

  server := httptest.NewServer(r)
  defer server.Close()

  for _, endpoint := range endpoints {
    res, err := http.Get(server.URL+endpoint)
    if err != nil {
      t.Fatal(err)
    }

  	if res.StatusCode != http.StatusOK {
  		t.Fatal(err)
  	}
  }
}

func TestNotFoundRoute(t *testing.T) {
  r := new(Router)
  r.HandleFunc(endpoints[0], testHandler1)
  r.HandleFunc(endpoints[1], testHandler2)

  server := httptest.NewServer(r)
  defer server.Close()

  res, err := http.Get(server.URL+"/test/handler/3")
    if err != nil {
      t.Fatal(err)
    }

  	if res.StatusCode != http.StatusNotFound {
  		t.Fatal(err)
  	}
}

func testHandler1(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "testHandler1 has been reached!")
}

func testHandler2(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "testHandler2 has been reached!")
}
