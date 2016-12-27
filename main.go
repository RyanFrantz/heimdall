package main

import (
    "log"
    "net/http"
    "github.com/RyanFrantz/heimdall/config"
    "github.com/julienschmidt/httprouter"
)

var cfg = config.GetConfig()

func main() {
    log.Printf("Starting Heimdall v%s", cfg.Heimdall.Version)
    router := httprouter.New()
    router.GET("/", indexResponse)
    router.GET("/ldap/group/:name", getLDAPGroup)
    router.GET("/ldap/user/:name", getLDAPUser)
    router.GET("/chef/client/:client", getChefClient)
    router.GET("/chef/group/:group", getChefGroup)

    // Start me up!
    log.Fatal(http.ListenAndServe(":8080", router))
}
