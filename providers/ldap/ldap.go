package ldap

import (
    "fmt"
    "log"
    "github.com/RyanFrantz/heimdall/config"
    "github.com/go-ldap/ldap"
)

var cfg = config.GetConfig()
var ldap_server     = cfg.Ldap.ServerAddress
var ldap_port       = cfg.Ldap.ServerPort
var ldap_base       = cfg.Ldap.SearchBase
var ldap_size_limit = cfg.Ldap.SearchSizeLimit
var ldap_time_limit = cfg.Ldap.SearchTimeLimit
var ldap_user       = cfg.Ldap.User
var ldap_password   = cfg.Ldap.Password

func ldapSearch(ldap_filter string, attributes []string) (searchResult *ldap.SearchResult) {
    l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldap_server, ldap_port))
    if err != nil {
        errorMessage := fmt.Sprintf("Failed to connect to host '%s' on port '%d' - %s", ldap_server, ldap_port, err)
        log.Fatal(errorMessage)
    } else {
        //log.Printf("Connect OK")
    }
    defer l.Close()

    bindRequest := ldap.NewSimpleBindRequest(ldap_user, ldap_password, nil)
    _, bindError := l.SimpleBind(bindRequest)

    if bindError != nil {
        log.Printf("ERROR: Cannot bind - " + bindError.Error())
    } else {
        //log.Printf("Bind OK")
    }

    searchRequest := ldap.NewSearchRequest(
        ldap_base,
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, ldap_size_limit, ldap_time_limit, false,
        ldap_filter,
        attributes,
        nil,
    )
    
    searchResult, searchError := l.Search(searchRequest)
    if searchError != nil {
        log.Fatal(searchError)
    }
    
    return searchResult
}

func GetLDAPGroup(name string, attributes []string) (searchResult *ldap.SearchResult) {
    var ldap_filter string
    if name == "all" {
        ldap_filter = "(&(objectClass=posixGroup))"
    } else {
        ldap_filter = fmt.Sprintf("(&(objectClass=posixGroup)(cn=%s))", name)
    }
    searchResult = ldapSearch(ldap_filter, attributes)
    return searchResult
}

func GetLDAPUser(name string, attributes []string) (searchResult *ldap.SearchResult) {
    var ldap_filter string
    if name == "all" {
        ldap_filter = "(&(objectClass=inetOrgPerson))"
    } else {
        ldap_filter = fmt.Sprintf("(&(objectClass=inetOrgPerson)(uid=%s))", name)
    }
    //attributes := []string{"dn", "cn", "memberUid"}
    searchResult = ldapSearch(ldap_filter, attributes)
    return searchResult
}
