{{- with secret "pki_int/cert/crl" -}}
{{- .Data.certificate }}
{{ end }}
{{- with secret "pki/cert/crl" -}}
{{- .Data.certificate }}
{{ end }}