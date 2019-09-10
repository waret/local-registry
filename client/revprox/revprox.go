package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
    "fmt"
)

func main() {
	var remoteSvc, localPort string
	if len(os.Args) == 3 {
		remoteSvc = os.Args[1]
		localPort = os.Args[2]
	} else {
		remoteSvc = os.Getenv("REMOTE_SVC")
		localPort = os.Getenv("LOCAL_PORT")
	}

    if remoteSvc == "" || localPort == "" {
        fmt.Println("empty parameters")
        os.Exit(1)
    }

	remote, err := url.Parse("http://" + remoteSvc)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":"+localPort, nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-Ben", "Rad")
		p.ServeHTTP(w, r)
	}
}

