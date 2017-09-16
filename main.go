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

	amech, err := client.Sys().ListAuth()
	if err != nil {
		log.Errorln(err)
	}
	for k, v := range amech {
		fmt.Println(k, tj(v))
	}

	// enable approle
	err = client.Sys().EnableAuthWithOptions("approle", &api.EnableAuthOptions{
		Type: "approle",
	})
	if err != nil {
		log.Errorln(err)
	}

	//approle, err := client.Logical().Write("auth/approle/role/testrole", map[string]interface{}{
	//
	//
	//})
	//if err != nil {
	//	log.Errorln(err)
	//}
}

func tj(v interface{}) string {
	cb, _ := json.MarshalIndent(v, "", "  ")
	return string(cb)
}

func jp(v interface{}) {
	cb, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(cb))
}
