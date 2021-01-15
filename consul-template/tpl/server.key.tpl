{{- with secret "pki_int/issue/verizon-dot-com" "ttl=30s" "common_name=demo.com" -}}
{{- .Data.private_key }}
{{ end }}