package main

import (
    "net/http"
    "io/ioutil"
    "strings"
    "log"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    agent := r.UserAgent()
    
    if path == "" { path = "main.html" }

    log.Println(strings.Split(r.RemoteAddr,":")[0] + " reqested", "/"+path, "\nagent: "+agent)
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
    http.Handle("/", new(MyHandler))
    http.ListenAndServe(":80", nil)
}
