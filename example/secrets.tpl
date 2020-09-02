full_json: {{ plugin "secretsmanager" "consul-template-app/example" }}
foo: {{ with plugin "secretsmanager" "consul-template-app/example"| parseJSON }}{{ base64Encode .foo }}{{ end }}
key: {{ with plugin "secretsmanager" "consul-template-app/example"| parseJSON }}{{ .key }}{{ end }}
sam: {{ with plugin "secretsmanager" "consul-template-app/example"| parseJSON }}{{ .sam }}{{ end }}
