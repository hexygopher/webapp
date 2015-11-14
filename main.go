package main

import (
	"github.com/stumpyfr/udger"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"flag"
	"fmt"
)

type MyHandler struct {
	u *udger.Udger
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := r.UserAgent()
	ua, err := this.u.Lookup(agent)
	path := r.URL.Path[1:]
	device := ua.Device.Name
	os := ua.OS.Family
	browser := ua.Browser.Name

	if err != nil {
		log.Fatal(err)
	}

	if path == "" {
		if device != "Smartphone" {
			path = "index.html"
		} else {
			path = "m/index.html"
		}
	}

	data, err := ioutil.ReadFile(string(path))
	
	if err == nil {
		w.Write(data)
	} else {
		data, _ = ioutil.ReadFile("404.html")
		w.WriteHeader(404)
		w.Write(data)
	}

	log.Println(strings.Split(r.RemoteAddr, ":")[0]+" reqested", "/"+path, "\ndevice:", device, "\nOS:", os, "\nBrowser:", browser, "\nUser-Agent: "+agent)

}
func main() {
	var err error
	port := flag.Int("p", 8080, "server port")
	flag.Parse()
	handler := new(MyHandler)
	handler.u, err = udger.New("udgerdb.dat")
	
	if err == nil {
		log.Println("UdgerDB loaded successfully")
	} else {
		log.Fatal(err)
	}
	
	http.Handle("/", handler)
	if err = http.ListenAndServe(fmt.Sprintf(":%d",*port), nil); err != nil{
		log.Fatal(err)
	}
}
