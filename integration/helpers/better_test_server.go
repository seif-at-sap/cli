package helpers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/onsi/gomega/ghttp"
)

type resp struct {
	status int
	body   []byte
}

var responses map[string]resp
var seenRoutes map[string]bool

func AddHandler(ser *ghttp.Server, method string, pathAndQuery string, status int, body []byte) {
	u, err := url.Parse(pathAndQuery)
	if err != nil {
		panic(err)
	}
	if len(responses) == 0 {
		responses = make(map[string]resp)
		seenRoutes = make(map[string]bool)
	}

	responses[key(method, u)] = resp{status, body}

	if !seenRoutes[method+u.Path] {
		ser.RouteToHandler(method, u.Path, func(w http.ResponseWriter, r *http.Request) {
			res, ok := responses[key(r.Method, r.URL)]
			if !ok {
				for k, v := range responses {
					fmt.Printf("For %s have status %d with %d bytes\n", k, v.status, len(v.body))
				}
				panic("no response for" + key(r.Method, r.URL))
			}
			println(res.status)
			println("**************************************")
			println(r.Method + ": " + r.URL.String())
			println("**************************************")
			w.WriteHeader(res.status)
			w.Write(res.body)
		})
		seenRoutes[method+u.Path] = true
	}
}

func key(method string, url *url.URL) string {
	return method + url.String()
}
