package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"log"
	"net/http"
)

var portNumber = flag.String("port", "9090", "port number.")
var basicAuthUser = flag.String("user", "", "basic auth user name")
var basicAuthPass = flag.String("pass", "", "basic auth user pass")

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	if *basicAuthUser != "" && *basicAuthPass != "" {
		log.Println("auth basic")
		auth.ProxyBasic(proxy, "yaproxy", func(user, pass string) bool {
			log.Println("auth basic do")
			return user == *basicAuthUser && pass == *basicAuthPass
		})
	}

	log.Println("listen:" + *portNumber)
	log.Fatal(http.ListenAndServe(":"+*portNumber, proxy))
}
