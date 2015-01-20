package route

import (
  "net/http"
  "sync"
  )


// inspired by Brad Fitzpatrick's idea:
// https://groups.google.com/forum/#!msg/golang-nuts/teSBtPvv1GQ/U12qA9N51uIJ
type Context struct {
  mutex sync.Mutex
  params map[*http.Request]map[string]string // URL parameters.
}

// setParams stores a map of URL paramters for a given request.
func (c *Context) setParams(req *http.Request, m map[string]string) {
  c.mutex.Lock()
  defer c.mutex.Unlock()

  if c.params == nil {
    c.params = make(map[*http.Request]map[string]string)
  }
  c.params[req] = m
}

// Get returns an URL parameter value for a given key for a given request.
func (c *Context) GetParam(req *http.Request, key string) string {
  c.mutex.Lock()
  defer c.mutex.Unlock()

  if c.params == nil {
    return ""
  }
  return c.params[req][key]
}

//  clear removes all the key/value pairs for a given request.
func (c *Context) clear(req *http.Request) {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  delete(c.params, req)
}
