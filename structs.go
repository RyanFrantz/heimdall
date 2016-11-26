package main

type Group struct {
    Dn              string  `json:"dn"`
    Name            string  `json:"name"`
    Gid             string  `json:"gid"`
    Members         []string `json:"members"`
}

type Groups []Group

type User struct {
    Dn              string  `json:"dn"`
    Uid             string  `json:"uid"`
    GivenName       string  `json:"givenName"`
    Sn              string  `json:"sn"`
    UidNumber       string  `json:"uidNumber"`
    GidNumber       string  `json:"gidNumber"`
    Description     string  `json:"description"`
    HomeDirectory   string  `json:"homeDirectory"`
}

type Users []User
