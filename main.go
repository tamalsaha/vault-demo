package main

import (
	"encoding/json"

	"github.com/appscode/log"
	"github.com/hashicorp/vault/api"
	"github.com/tamalsaha/go-oneliners"
)

// https://www.vaultproject.io/docs/concepts/response-wrapping.html#response-wrapping-token-creation
// Response wrapping is per-request and is triggered by providing to Vault the desired TTL for a response-wrapping token for that request. This is set by the client using the X-Vault-Wrap-TTL header and can be either an integer number of seconds or a string duration of seconds (15s), minutes (20m), or hours (25h)

func main() {
	cfg := api.DefaultConfig()
	oneliners.FILE(tj(cfg))

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Errorln(err)
	}
	if err != nil {
		log.Errorln(err)
	}

	//s, err := client.Auth().Token().LookupSelf()
	//oneliners.FILE(tj(s))

	enableAppRole(client)

	roles, err := client.Logical().List("/auth/approle/role")
	if err != nil {
		log.Errorln(err)
	}
	oneliners.FILE(tj(roles))

	r2, err := client.Logical().Write("auth/approle/role/testrole", map[string]interface{}{
		"secret_id_ttl":      "10m",
		"token_num_uses":     "10",
		"token_ttl":          "2m",
		"token_max_ttl":      "30m",
		"policies":           []string{"dev-policy", "test-policy"},
		"secret_id_num_uses": 80,
	})
	oneliners.FILE(tj(r2), err)

	r3, err := client.Logical().Read("auth/approle/role/testrole/role-id")
	oneliners.FILE(tj(r3.Data["role_id"]), err)

	//client.SetWrappingLookupFunc(func(operation, path string) string {
	//	return "10m"
	//})

	r4, err := client.Logical().Write("auth/approle/role/testrole/secret-id", map[string]interface{}{
			"host_ip":   "1.2.3.4",
	})
	oneliners.FILE(tj(r4), err)
	oneliners.FILE(r4.Data["secret_id"], "|", r4.Data["secret_id_accessor"])

	r5, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   r3.Data["role_id"],
		"secret_id": r4.Data["secret_id"],
	})
	oneliners.FILE(tj(r5), err)

	//tcr := &api.TokenCreateRequest{
	//	Policies: []string{"myrole"},
	//	Metadata: map[string]string{
	//		"host_ip":   "1.2.3.4",
	//		//"namespace": pod.Metadata.Namespace,
	//		//"pod_ip":    pod.Status.PodIP,
	//		//"pod_name":  pod.Metadata.Name,
	//		//"pod_uid":   pod.Metadata.Uid,
	//	},
	//	DisplayName: "pod.Metadata.Name",
	//	Period:      "100h",
	//	NoParent:    true,
	//	TTL:         "100h",
	//}
	//r6, err := client.Auth().Token().Create(tcr)
	//if err != nil {
	//	log.Errorln(err)
	//}
	//oneliners.FILE(tj(r6), err)

	r7, err := client.Logical().Write( "/auth/approle/role/testrole/secret-id/destroy", map[string]interface{}{
		"secret_id": r4.Data["secret_id"],
	})
	if err != nil {
		log.Errorln(err)
	}
	oneliners.FILE(tj(r7), err)

	r8, err := client.Auth().Token().RenewTokenAsSelf(r5.Auth.ClientToken, 0)
	if err != nil {
		log.Errorln(err)
	}
	oneliners.FILE(tj(r8), err)

	// renew token
	// https://github.com/Boostport/kubernetes-vault/blob/5175dd51d5c72efd1b5c18303bb3940b98e13983/cmd/controller/client/vault.go#L371
}

func tj(v interface{}) string {
	cb, _ := json.MarshalIndent(v, "", "  ")
	return string(cb)
}

func enableAppRole(client *api.Client) {
	// list enabled auth mechanism
	amech, err := client.Sys().ListAuth()
	if err != nil {
		log.Errorln(err)
	}
	for k, v := range amech {
		if k == "approle/" && v.Type == "approle" {
			oneliners.FILE("approle enabled _________________")
			return
		}
	}

	// $ vault auth-enable approle
	err = client.Sys().EnableAuthWithOptions("approle", &api.EnableAuthOptions{
		Type: "approle",
	})
	if err != nil {
		log.Errorln(err)
	}
}
