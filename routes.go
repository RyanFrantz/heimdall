package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/RyanFrantz/heimdall/providers/ldap"
    "github.com/julienschmidt/httprouter"
)

func indexResponse (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "I am Heimdall\n")
}

func getGroup (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    w.Header().Set("Cache-Control", "max-age=3600")

    attributes := []string{"cn", "gidNumber", "memberUid"}
    results := ldap.GetLDAPGroup(name, attributes)
    var groups Groups
    for _, entry := range results.Entries {
        groupName := entry.GetAttributeValue("cn")
        gid := entry.GetAttributeValue("gidNumber")
        members := entry.GetAttributeValues("memberUid")
        group := Group{Dn: entry.DN, Name: groupName, Gid: gid, Members: members}
        groups = append(groups, group)
    }
    json.NewEncoder(w).Encode(groups)
}

func getUser (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    attributes := []string{"uid", "cn", "uidNumber", "gidNumber", "givenName", "sn", "description", "homeDirectory"}
    results := ldap.GetLDAPUser(name, attributes)
    var users Users
    for _, entry := range results.Entries {
        uid := entry.GetAttributeValue("uid")
        uidNumber := entry.GetAttributeValue("uidNumber")
        gidNumber := entry.GetAttributeValue("gidNumber")
        givenName := entry.GetAttributeValue("givenName")
        sn := entry.GetAttributeValue("sn")
        description := entry.GetAttributeValue("description")
        homeDirectory := entry.GetAttributeValue("homeDirectory")
        user := User{
            Dn: entry.DN,
            Uid: uid,
            UidNumber: uidNumber,
            GidNumber: gidNumber,
            GivenName: givenName,
            Sn: sn,
            Description: description,
            HomeDirectory: homeDirectory}
        users = append(users, user)
    }
    w.Header().Set("Cache-Control", "max-age=3600")
    json.NewEncoder(w).Encode(users)
}
