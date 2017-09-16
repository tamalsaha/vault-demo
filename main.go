package main

import (
	"encoding/json"
	"fmt"
	"github.com/appscode/log"
	"github.com/hashicorp/vault/api"
	"github.com/tamalsaha/go-oneliners"
)

func main() {
	cfg := api.DefaultConfig()
	oneliners.FILE(tj(cfg))

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Errorln(err)
	}
	s, err := client.Auth().Token().LookupSelf()
	oneliners.FILE(tj(s))

	// list enabled auth mechanism
	amech, err := client.Sys().ListAuth()
	if err != nil {
		log.Errorln(err)
	}
	for k, v := range amech {
		fmt.Println(k, tj(v))
	}

	/*
	// $ vault auth-enable approle
	err = client.Sys().EnableAuthWithOptions("approle", &api.EnableAuthOptions{
		Type: "approle",
	})
	if err != nil {
		log.Errorln(err)
	}
	*/

	roles, err := client.Logical().List("/auth/approle/role")
	if err != nil {
		log.Errorln(err)
	}
	oneliners.FILE(tj(roles))

	approle, err := client.Logical().Write("auth/approle/role/testrole", map[string]interface{}{
		"secret_id_ttl":      "10m",
		"token_num_uses":     "10",
		"token_ttl":          "20m",
		"token_max_ttl":      "30m",
		"secret_id_num_uses": 40,
	})
	if err != nil {
		log.Errorln(err)
	}
	oneliners.FILE(tj(approle))
}

func tj(v interface{}) string {
	cb, _ := json.MarshalIndent(v, "", "  ")
	return string(cb)
}
