package cmd

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "strconv"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
}

func handleDisk(w http.ResponseWriter, r *http.Request) {
    _, err := ioutil.ReadFile(filepath)
  	if err != nil {
  	       http.Error(w, err.Error(), http.StatusInternalServerError)
           return
  	}
    fmt.Fprintf(w, "ok")
}

func handleRemote(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get(remoteurl)
    if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
    }
    defer resp.Body.Close()
    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
    }
    fmt.Fprintf(w, "ok")
}

func runServer() {
    http.HandleFunc("/healthz", handleHealth)
    http.HandleFunc("/disk", handleDisk)
    http.HandleFunc("/remote", handleRemote)

    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
