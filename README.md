# vault-demo
Learning Vault api

```console
# start server
$ vault server -dev -dev-root-token-id=3e4a5ba1-kube-422b-d1db-844979cab098

# export vault config
$ export VAULT_ADDR='http://127.0.0.1:8200'
$ export VAULT_TOKEN='3e4a5ba1-kube-422b-d1db-844979cab098'

# run demo
$ go run main.go
```
