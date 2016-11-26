package main

import (
    "log"
    "net/http"
    "github.com/RyanFrantz/heimdall/config"
    "github.com/julienschmidt/httprouter"
)

// Global heimdall config variable.
var hconfig = config.ReadConfig()

func main() {
    //http.HandleFunc("/", indexResponse) // Assumes DefaultServeMux
    // nil here effectively tells net/http we'll use the DefaultServeMux multiplexer for routing requests to functions.
    //log.Fatal(http.ListenAndServe(":8080", nil))

    log.Printf("Starting Heimdall v%s", hconfig.Heimdall.Version)
    router := httprouter.New()
    router.GET("/", indexResponse)
    router.GET("/group/:name", getGroup)
    router.GET("/user/:name", getUser)

    log.Fatal(http.ListenAndServe(":8080", router))

}
