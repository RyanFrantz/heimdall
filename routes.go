package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/RyanFrantz/heimdall/plugins/chef"
    "github.com/RyanFrantz/heimdall/plugins/ldap"
    "github.com/julienschmidt/httprouter"
)

func indexResponse (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "I am Heimdall\n")
}

func getChefClient (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("client")
    w.Header().Set("Cache-Control", "max-age=3600")

    results := chef.GetClient(name)
    json.NewEncoder(w).Encode(results)
}

func getChefGroup (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    group_name := ps.ByName("group")
    w.Header().Set("Cache-Control", "max-age=3600")

    results := chef.GetGroup(group_name)
    json.NewEncoder(w).Encode(results)
}

func getLDAPGroup (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    w.Header().Set("Cache-Control", "max-age=3600")

    attributes := []string{"cn", "gidNumber", "memberUid"}
    results := ldap.GetGroup(name, attributes)
    var groups LdapGroups
    for _, entry := range results.Entries {
        groupName := entry.GetAttributeValue("cn")
        gid := entry.GetAttributeValue("gidNumber")
        members := entry.GetAttributeValues("memberUid")
        group := LdapGroup{Dn: entry.DN, Name: groupName, Gid: gid, Members: members}
        groups = append(groups, group)
    }
    json.NewEncoder(w).Encode(groups)
}

func getLDAPUser (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    attributes := []string{"uid", "cn", "uidNumber", "gidNumber", "givenName", "sn", "description", "homeDirectory"}
    results := ldap.GetUser(name, attributes)
    var users LdapUsers
    for _, entry := range results.Entries {
        uid := entry.GetAttributeValue("uid")
        uidNumber := entry.GetAttributeValue("uidNumber")
        gidNumber := entry.GetAttributeValue("gidNumber")
        givenName := entry.GetAttributeValue("givenName")
        sn := entry.GetAttributeValue("sn")
        description := entry.GetAttributeValue("description")
        homeDirectory := entry.GetAttributeValue("homeDirectory")
        user := LdapUser{
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
