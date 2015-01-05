package route

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var endpoints = []string{
	"/test/handler/1/?",
	"/test/handler/2/?",
}

func TestRouter(t *testing.T) {
	r := new(Router)
	r.HandleFunc(endpoints[0], handler1)
	r.HandleFunc(endpoints[1], handler2)

	server := httptest.NewServer(r)
	defer server.Close()

	var urls = []string{
		"/test/handler/1",
		"/test/handler/2",
		"/test/handler/1/",
		"/test/handler/2/",
	}

	for _, u := range urls {
		res, err := http.Get(server.URL + u)
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
	r.HandleFunc(endpoints[0], handler1)
	r.HandleFunc(endpoints[1], handler2)

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
	r.HandleFunc("/test/handler/[0-9]+/hello/?", handlerHello)

	urls := []string{
		"/test/handler/1/hello",
		"/test/handler/2/hello/",
		"/test/handler/2123/hello/",
		"/test/handler/2123/hello",
	}

	server := httptest.NewServer(r)
	defer server.Close()

	for _, u := range urls {
		res, err := http.Get(server.URL + u)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Fatal(res)
		}
	}

}

func TestNotFoundRegexpRoute(t *testing.T) {
	r := new(Router)
	r.HandleFunc("/test/handler/[0-9]/hello", handlerHello)

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

func TestServeStaticResources(t *testing.T) {
	rootTestingPath := "/temp_TestServeStaticResources"
	cssTestingPath := rootTestingPath+"/css"
	jsTestingPath := rootTestingPath+"/js"

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&rootTestingPath)
	r.AddStaticResource(&cssTestingPath)
	r.AddStaticResource(&jsTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func TestServeTwoLevelStaticResources(t *testing.T) {
	rootTestingPath := "/temp_TestServeStaticResources"
	cssTestingPath := rootTestingPath+"/css"

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&cssTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func TestServeNonExistingStaticResources(t *testing.T) {
	rootTestingPath := "/temp_TestServeStaticResources"
	jsTestingPath := rootTestingPath+"/js"

	createTestingData(rootTestingPath)

	r := new(Router)
	r.AddStaticResource(&jsTestingPath)

	server := httptest.NewServer(r)
	defer server.Close()

	res, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Fatal(res)
	}

	cleanTestingData(rootTestingPath)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler1 has been reached!")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandler2 has been reached!")
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testHandlerHello has been reached!")
}

func createTestingData(rootTestPath string) {
	src, err := os.Stat(rootTestPath)
	if err != nil || !src.IsDir() {
		os.Mkdir(rootTestPath, 0777)
	}

	src, err = os.Stat(rootTestPath+"/index.html")
	if err != nil || src.IsDir() {
		os.Create(rootTestPath+"/index.html")
	}

	cssTestPath := rootTestPath+"/css"

	src, err = os.Stat(cssTestPath)
	if err != nil || !src.IsDir() {
		os.Mkdir(cssTestPath, 0777)
	}

	src, err = os.Stat(cssTestPath+"/main.css")
	if err != nil || src.IsDir() {
		os.Create(cssTestPath+"/main.css")
	}
}

func cleanTestingData(nameTest string) {
	src, err := os.Stat(nameTest)
	if err == nil && src.IsDir() {
		os.RemoveAll(nameTest)
	}
}
