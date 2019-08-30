# consul-template-plugin-secretsmanager

Plugin for consul-template to retrieve and decrypt from AWS Secrets Manager

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
{{ plugin "aws-secretsmanager" "path/to/secrets" }}
```

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
