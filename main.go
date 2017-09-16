package main

import (
	"github.com/hashicorp/vault/api"
	"encoding/json"
	"fmt"
	"github.com/appscode/log"
)

func main() {
	cfg := api.DefaultConfig()
	cb, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(cb))

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
