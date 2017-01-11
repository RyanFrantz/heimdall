package main

type LdapGroup struct {
    Dn              string  `json:"dn"`
    Name            string  `json:"name"`
    Gid             string  `json:"gid"`
    Members         []string `json:"members"`
}

type LdapGroups []LdapGroup

// Create a map whose key is a string and whose value is a slice of strings.
//type LdapGroupMap {}

type LdapUser struct {
    Dn              string  `json:"dn"`
    Uid             string  `json:"uid"`
    GivenName       string  `json:"givenName"`
    Sn              string  `json:"sn"`
    UidNumber       string  `json:"uidNumber"`
    GidNumber       string  `json:"gidNumber"`
    Description     string  `json:"description"`
    HomeDirectory   string  `json:"homeDirectory"`
}

type LdapUsers []LdapUser
