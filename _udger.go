package main

import (
    "net/http"
    "io/ioutil"
    "strings"
    "log"
    "github.com/stumpyfr/udger"
)

type MyHandler struct {
    u *udger.Udger
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    agent := r.UserAgent()
    
    if path == "" { path = "main.html" }

    log.Println(strings.Split(r.RemoteAddr,":")[0] + " reqested", "/"+path, "\nagent: "+agent)
	ua, err := this.u.Lookup(agent)
		if err != nil {
			log.Println("error: ", err)
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
    h:=new(MyHandler)
    h.u,_=udger.New("udgerdb.dat")
    http.Handle("/",h) 
    http.ListenAndServe(":8080", nil)
}
