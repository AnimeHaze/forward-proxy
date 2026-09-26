package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	httpproxy "github.com/orkunkaraduman/go-httpproxy"
)

func OnError(ctx *httpproxy.Context, where string,
	err *httpproxy.Error, opErr error) {
	log.Printf("ERR: %s: %s [%s]", where, err, opErr)
}

func OnRequest(ctx *httpproxy.Context, req *http.Request) (resp *http.Response) {
	log.Printf("INFO: Proxy: %s %s", req.Method, req.URL.String())
	return nil
}

func OnConnect(ctx *httpproxy.Context, host string) (httpproxy.ConnectAction, string) {
	log.Printf("HTTPS CONNECT %s", host)

	return httpproxy.ConnectProxy, ""
}

func OnResponse(ctx *httpproxy.Context, req *http.Request,
	resp *http.Response) {
	resp.Header.Add("Via", "go-httpproxy")
}

func main() {
	port := flag.String("port", "8033", "Port to listen on")
	host := flag.String("host", "", "Host to bind to (optional)")
	verbose := flag.Bool("verbose", true, "Enable verbose logging")
	version := flag.Bool("version", false, "Print version and exit")

	flag.Parse()
	
	if *version {
	  fmt.Println("2.0.0")
	  return
  }

	prx, err := httpproxy.NewProxy()
	if err != nil {
		log.Fatal(err)
	}

  if *verbose {
	  prx.OnError = OnError
	  prx.OnConnect = OnConnect
	  prx.OnRequest = OnRequest
	  prx.OnResponse = OnResponse
  }

	addr := fmt.Sprintf(":%s", *port)

	if *host != "" {
		addr = fmt.Sprintf("%s:%s", *host, *port)
	}

	fmt.Printf("Proxy server starting on %s\n", addr)
	fmt.Println("OK_SUCCESS")

	if err := http.ListenAndServe(addr, prx); err != nil {
		log.Fatal(err)
	}
}
