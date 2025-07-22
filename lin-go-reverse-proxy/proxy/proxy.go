package proxy

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "sync/atomic"
)
type Proxy struct {
    targets []*url.URL
    counter uint64
}
func NewProxy(targets []string) *Proxy {
    parsed := []*url.URL{}
    for _, t := range targets {
        u, err := url.Parse(t)
        if err != nil {
            log.Fatalf("Invalid URL %s: %v", t, err)
        }
        parsed = append(parsed, u)
    }
    return &Proxy{targets: parsed}
}
func (p *Proxy) getNextTarget() *url.URL {
    i := atomic.AddUint64(&p.counter, 1)
    return p.targets[(int(i)-1)%len(p.targets)]
}
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    target := p.getNextTarget()
    proxy := httputil.NewSingleHostReverseProxy(target)
    r.Host = target.Host
    log.Printf("\u2192 Forwarding request: %s %s --> %s", r.Method, r.URL.Path, target)
    proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
        log.Printf("!! Proxy error for %s: %v", r.URL.Path, err)
        http.Error(w, "Reverse proxy error", http.StatusBadGateway)
    }
    proxy.ServeHTTP(w, r)
}