package main

import (
	"encoding/json"
	"fmt"
	"github.com/appscode/log"
	"github.com/hashicorp/vault/api"
)

func main() {
	cfg := api.DefaultConfig()
	jp(cfg)

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Errorln(err)
	}
	s, err := client.Auth().Token().LookupSelf()
	jp(s)
}

func jp(v interface{}) {
	cb, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(cb))
}
