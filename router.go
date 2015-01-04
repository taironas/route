package route

import (
	"net/http"
	"regexp"
)

// inspired by the following sources with some small changes:
// http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc
// https://github.com/raymi/quickerreference
type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	routes       []*route  // array of routes with a tuple (pattern, handler)
	staticRoutes []*string // array of static routes
}

// Handle registers the handler for the given pattern in the router.
func (r *Router) Handle(strPattern string, handler http.Handler) {
	// encapsulate string pattern with start and end constraints.
	pattern := regexp.MustCompile("^" + strPattern + "$")
	r.routes = append(r.routes, &route{pattern, handler})
}

// HandleFunc registers the handler function for the given pattern in the router.
func (r *Router) HandleFunc(strPattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.Handle(strPattern, http.HandlerFunc(handler))
}

// ServeHTTP looks for a matching route among the routes. Returns 404 if no match is found.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for _, route := range r.routes {
		if route.pattern.MatchString(req.URL.Path) {
			route.handler.ServeHTTP(w, req)
			return
		}
	}

	// route not found. check if it is a static ressource.
	for _, sr := range r.staticRoutes {
		dir := http.Dir(*sr)
		if _, err := dir.Open(req.URL.Path); err == nil {
			// Could open file, set static route and call ServeHTTP again.
			r.Handle(req.URL.Path, http.FileServer(dir))
			r.ServeHTTP(w, req)
			return
		}
	}

	// no pattern matched; send 404 response
	http.NotFound(w, req)
}

// AddStaticRoute adds a route value to an array of static routes.
// Use this is you want to serve a static directory and it's sub directories.
func (r *Router) AddStaticRoute(route *string) {
	r.staticRoutes = append(r.staticRoutes, route)
}
