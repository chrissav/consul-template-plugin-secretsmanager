<p align="center">
  <a href="https://hellofresh.com">
    <img width="120" src="https://www.hellofresh.de/images/hellofresh/press/HelloFresh_Logo.png">
  </a>
</p>

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

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
