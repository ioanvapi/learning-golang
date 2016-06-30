package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

// Prox our RerverseProxy object
type Prox struct {
	// target url of reverse proxy
	target        *url.URL
	// instance of Go ReverseProxy that will do the job for us
	proxy         *httputil.ReverseProxy
	// add some route patterns with regexp
	routePatterns []*regexp.Regexp
}

// small factory
func New(target string) *Prox {
	// you should handle error on parsing
	url, _ := url.Parse(target)

	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")

	if p.routePatterns == nil || p.parseWhiteList(r) {
		// call to magic method from ReverseProxy object
		p.proxy.ServeHTTP(w, r)
	}
}

func (p *Prox) parseWhiteList(r *http.Request) bool {
	for _, regexp := range p.routePatterns {
		fmt.Println(r.URL.Path)
		if regexp.MatchString(r.URL.Path) {
			return true
		}
	}
	fmt.Println("Not accepted routes %x", r.URL.Path)
	return false
}

func main() {
	const (
		defaultPort             = ":80"
		defaultPortUsage        = "default server port, ':80', ':8080'..."
		defaultTarget           = "http://127.0.0.1:8080"
		defaultTargetUsage      = "default redirect url, 'http://127.0.0.1:8080'"
		defaultWhiteRoutes      = `^\/$|[\w|/]*.js|/path|/path2`
		defaultWhiteRoutesUsage = "list of white route as regexp, '/path1*,/path2*...."
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	url := flag.String("url", defaultTarget, defaultTargetUsage)
	routesRegexp := flag.String("routes", defaultWhiteRoutes, defaultWhiteRoutesUsage)

	flag.Parse()

	fmt.Println("server will run on : %s", *port)
	fmt.Println("redirecting to :%s", *url)
	fmt.Println("accepted routes :%s", *routesRegexp)

	reg, _ := regexp.Compile(*routesRegexp)
	routes := []*regexp.Regexp{reg}

	// proxy
	proxy := New(*url)
	proxy.routePatterns = routes

	// server
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(*port, nil)
}
