# consul-template-plugin-ssm

Plugin for consul-template to retrieve and decrypt from AWS SSM and KMS.

## Building
Run make from the root directory:

`make deps`

`make build`

The built binary is placed in the `dist/` directory

## Running tests

`make test`

## Usage

`export AWS_REGION=<AWS REGION>`

Place the binary in the PATH (or use the full path), inside the consul-template use:
```
{{ plugin: "ssm" <parameter name> }}
```