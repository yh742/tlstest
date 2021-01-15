vault {
  address = "http://132.197.249.120:8200"
  token="s.I2dgamZ508RVxiIuYXPBaGEF"  #also can be under VAULT_TOKEN env variable
  unwrap_token = false
  renew_token = false
}
syslog {
  enabled = true
  facility = "LOCAL5"
}
template {
    source="./tpl/server.crt.tpl"
    destination="./secrets/client.crt"
    #Optional Command after certificate renewal
    command = "mv ./secrets/client.crt ../sds/secrets/client.crt"
}
template {
    source="./tpl/server.key.tpl"
    destination="./secrets/client.key"
    #Optional Command after certificate renewal
    command = "mv ./secrets/client.key ../sds/secrets/client.key"
}
template {
    source="./tpl/ca.crt.tpl"
    destination="./secrets/ca.crt"
    #Optional Command after certificate renewal
    command = "mv ./secrets/ca.crt ../sds/secrets/ca.crt"
}
template {
    source="./tpl/crl.tpl"
    destination="./secrets/crl.pem"
    #Optional Command after certificate renewal
    command = "mv ./secrets/crl.pem ../sds/secrets/crl.pem"
}