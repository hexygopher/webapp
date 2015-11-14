package main

import (
	"github.com/stumpyfr/udger"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {
	u *udger.Udger
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "main.html"
	}
	agent := r.UserAgent()
	log.Println(strings.Split(r.RemoteAddr, ":")[0]+" reqested", "/"+path)
	ua, err := this.u.Lookup(agent)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", ua)
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
		w.Write(data)
	} else {
		data, _ = ioutil.ReadFile("404.html")
		w.WriteHeader(404)
		w.Write(data)
	}
}
func main() {
	handler := new(MyHandler)
	var err error
	handler.u, err = udger.New("udgerdb.dat")
	if err == nil {
		log.Println("UdgerDB loaded successfully")
	} else {
		log.Fatal(err)
	}
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
