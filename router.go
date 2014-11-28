package route

import (
	"net/http"
	"regexp"
)

// inspired by the following sources with some small changes:
//http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc
//https://github.com/raymi/quickerreference
type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	routes []*route
}

// HandleFunc registers the handler function for the given pattern in the router.
func (r *Router) HandleFunc(strPattern string, handler func(http.ResponseWriter, *http.Request)) {
	// encapsulate string pattern with start and end constraints.
	pattern := regexp.MustCompile("^" + strPattern + "$")
	r.routes = append(r.routes, &route{pattern, http.HandlerFunc(handler)})
}

// ServeHTTP looks for a matching route among the routes. Returns 404 if no match is found.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.pattern.MatchString(req.URL.Path) {
			route.handler.ServeHTTP(w, req)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, req)
}
