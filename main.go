package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "strings"
    
    "github.com/elazarl/goproxy"
)

func main() {
    // Define a port flag with default value
    port := flag.String("port", "8033", "Port to listen on")
    host := flag.String("host", "", "Host to bind to (optional)")
    verbose := flag.Bool("verbose", true, "Enable verbose logging")

    flag.Parse()
    
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = *verbose

    var addr string
    if *host != "" {
        addr = fmt.Sprintf("%s:%s", *host, *port)
    } else {
        addr = fmt.Sprintf(":%s", *port)
    }
    
    fmt.Printf("Proxy server starting on %s\n", addr)
    fmt.Printf("OK_SUCCESS\n", addr)
    log.Fatal(http.ListenAndServe(addr, proxy))
    fmt.Printf("ERROR_FATAL\n", addr)
}
