{{- with secret "pki_int/issue/verizon-dot-com" "ttl=30s" "common_name=demo.com" -}}
{{- .Data.certificate }}
{{ end }}
{{- with secret "pki_int/issue/verizon-dot-com" "ttl=30s" "common_name=demo.com" -}}
{{- .Data.issuing_ca }}
{{ end }}