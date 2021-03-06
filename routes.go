package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/RyanFrantz/heimdall/plugins/souschef"
    "github.com/RyanFrantz/heimdall/plugins/ldap"
    "github.com/julienschmidt/httprouter"
    "github.com/RyanFrantz/chef"
)

func indexResponse (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "I am Heimdall\n")
}

func getChefClient (name string) (chef_client chef.ApiClient) {
    chef_client = souschef.GetClient(name)
    return chef_client
}

// Route handler for looking up Chef client objects.
func getChefClientForRequest (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("client")
    w.Header().Set("Cache-Control", "max-age=3600")
    results := getChefClient(name)
    json.NewEncoder(w).Encode(results)
}

func getChefGroup (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    group_name := ps.ByName("group")
    w.Header().Set("Cache-Control", "max-age=3600")

    results := souschef.GetGroup(group_name)
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

// Look up an LDAP user object.
// Expects the name of an LDAP user or the special keyword 'all' to look up all
// LDAP users.
// Returns an array of type LdapUsers.
func getLdapUser (name string) (users LdapUsers) {
    attributes := []string{"uid", "cn", "uidNumber", "gidNumber", "givenName", "sn", "description", "homeDirectory"}
    results := ldap.GetUser(name, attributes)
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
    return users
}

// Route handler for looking up LDAP user objects.
func getLdapUserForRequest (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    users := getLdapUser(name)
    w.Header().Set("Cache-Control", "max-age=3600")
    json.NewEncoder(w).Encode(users)
}

func getUserReport (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    users := getLdapUser(name)
    chef_client := getChefClient(name)
    w.Header().Set("Cache-Control", "max-age=3600")
    json.NewEncoder(w).Encode(users)
    json.NewEncoder(w).Encode(chef_client)
}
