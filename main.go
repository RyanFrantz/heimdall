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
    log.Printf("Starting Heimdall v%s", hconfig.Heimdall.Version)
    router := httprouter.New()
    router.GET("/", indexResponse)
    router.GET("/ldap/group/:name", getGroup)
    router.GET("/ldap/user/:name", getUser)

    // Start me up!
    log.Fatal(http.ListenAndServe(":8080", router))
}
