# consul-template-plugin-secretsmanager

Plugin for consul-template to retrieve and decrypt from AWS Secrets Manager

## Building
Run make from the root directory:

`make deps`

`make build`

The built binary is placed in the `dist/` directory

## Installing

To install with `go`:
`go get -v github.com/chrissav/consul-template-plugin-secretsmanager`

Rename, move, or symlink the binary to `secretsmanager` in your path.

For example (this assumes your $GOPATH is in your $PATH):
```
ln -s $(which consul-template-plugin-secretsmanager) ~/.go/bin/secretsmanager
```

Confirm `secretsmanager` is in your path:
```
which secretsmanager
```

## Running tests

`make test`

## Usage

`export AWS_REGION=<AWS REGION>`

Place the binary in the PATH (or use the full path), inside the consul-template use:
```
{{ plugin "secretsmanager" "path/to/secrets" }}
```

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
