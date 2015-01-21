package route

import (
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)

	var context Context

	foo := context.GetParam(r, "foo")
	if len(foo) > 0 {
		t.Fatal("Expected an empty string and got", foo)
	}

	params := make(map[string]string)
	params["foo"] = "johndoe"
	params["foo2"] = "42"

	context.setParams(r, params)

	if len(context.params[r]) != len(params) {
		t.Fatal("Params map should contained", len(params), "elements. context params length =", len(context.params[r]))
	}

	id := context.GetParam(r, "foo2")
	if id != params["foo2"] {
		t.Fatal("Expected", params["foo2"], "and got", id)
	}

	context.clear(r)
	if len(context.params) != 0 {
		t.Fatal("Params map should be empty")
	}
}
