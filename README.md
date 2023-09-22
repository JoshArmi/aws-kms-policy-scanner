# AWS KMS Policy Scanner

This project contains the initial work to build a tool capable of identifying common anti-patterns in AWS KMS Key resource policies.

For more info, see this [blog post](https://medium.com/@josh.armitage/the-two-common-aws-kms-anti-patterns-98ad7b7e43a3)

## Usage

1. Run `make`
2. Run `assume <your-role> --exec ./main`

In your console it will print out all the identified keys

## Known Issues

1. Currently doesn't differentiate between AWS Managed Keys and Customer Managed Keys.
2. `assume` doesn't correctly populate env vars, forcing of wrapping execution in `assume > your-role> --exec`
