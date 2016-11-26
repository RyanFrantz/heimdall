package main

// Initial inspiration from https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/

import (
    "fmt"
    "log"
    "os"
    "github.com/BurntSushi/toml"
    //"github.com/spf13/viper" // Perhaps one day when we need more features.
)

type HeimdallConfig struct {
    Version     string
    ServerPort  int
}

type LdapConfig struct {
    ServerAddress   string
    ServerPort      int
    SearchBase      string
    SearchSizeLimit int
    SearchTimeLimit int
    User            string
    Password        string
}

type Config struct {
    Heimdall    HeimdallConfig
    Ldap        LdapConfig
}

func ReadConfig() Config {
    configfile := "./heimdall.conf"
    _, statErr := os.Stat(configfile)
    if statErr != nil {
        statMsg := fmt.Sprintf("ERROR: Unable to stat config file '%s': %v\n", configfile, statErr.Error())
        log.Fatal(statMsg)
    }

    var config Config
    if _, decodeErr := toml.DecodeFile(configfile, &config); decodeErr != nil {
        decodeMsg := fmt.Sprintf("ERROR: Cannot parse '%s': %v\n", configfile, decodeErr.Error())
        log.Fatal(decodeMsg)
    }
    return config
}
