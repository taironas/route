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
		res, err := http.Get(server.URL + endpoint)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatal(res)
		}
	}
}

func TestNotFoundRoute(t *testing.T) {
	r := new(Router)
	r.HandleFunc(endpoints[0], testHandler1)
	r.HandleFunc(endpoints[1], testHandler2)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/3")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}
}

func TestFoundRegexpRoute(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/[0-9]/hello", testHandlerHello)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/1/hello")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}
}

func TestNotFoundRegexpRoute(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/[0-9]/hello", testHandlerHello)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/test/handler/a/hello")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}
}

func testHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler1 has been reached!")
}

func testHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler2 has been reached!")
}

func testHandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandlerHello has been reached!")
}
