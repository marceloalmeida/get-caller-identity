# get-caller-identity

A simple Go CLI tool to retrieve the AWS caller identity using the AWS SDK for Go v2.

## Overview
This tool calls the AWS STS `GetCallerIdentity` API to display the current AWS user's UserId, Account, and ARN. It is useful for debugging AWS credentials and identity issues.

## Prerequisites
- Go 1.18 or newer
- AWS credentials configured (via environment variables, AWS config files, or other supported methods)

## Installation
Clone the repository and build the binary:

```sh
git clone https://github.com/marceloalmeida/get-caller-identity.git
cd get-caller-identity
go build -o get-caller-identity main.go
```

## Usage
Run the tool:

```sh
./get-caller-identity
```

You should see output similar to:

```
UserId: AIDAEXAMPLEID
Account: 123456789012
ARN: arn:aws:iam::123456789012:user/example-user
```

## How it works
- Loads AWS configuration using the default provider chain
- Creates an STS client
- Calls `GetCallerIdentity`
- Prints UserId, Account, and ARN

## License
See [LICENSE](LICENSE) for details.
