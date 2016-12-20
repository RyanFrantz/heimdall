package chef

import (
    "log"
    "io/ioutil"
    "github.com/RyanFrantz/heimdall/config"
    "github.com/go-chef/chef"
)

var cfg = config.GetConfig()
var chef_server = cfg.Chef.ServerAddress

func GetClient(name string) (client chef.ApiClient) {
    key_file := "heimdall.pem"
    client_key, err := ioutil.ReadFile(key_file)
    if err != nil {
        log.Printf("Unable to read Chef client key '%s': %s\n", key_file, err)
    }

    // Create a client object.
    heimdall_client, new_client_err := chef.NewClient(&chef.Config{
        Name: "heimdall",
        Key: string(client_key),
        BaseURL: chef_server,
    })

    if new_client_err != nil {
        log.Printf("Failed to create Chef client object: %s", new_client_err)
    }

    client, get_client_err := heimdall_client.Clients.Get(name)
    if get_client_err != nil {
        log.Printf("Failed to get info for client '%s': %s", name, get_client_err)
    }

    // The client is already JSON.
    return client
}