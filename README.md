# Keygen

This repository provide a simple console application for secure keys generation to use with Schnorr signatures,
and FROST protocol participants.

`keygen` randomly generated key pair and prints it to stdout or specified file, depending on the provided arguments
(please see below or help command for details).

## Motivation

Schnorr signatures standard is described in
[Bitcoin Improvement Proposal #0340](https://github.com/bitcoin/bips/blob/master/bip-0340.mediawiki) requires
public keys being with valid odd Y coordinate.   
So this project is aimed to simplify keys generation for users convenience.

## Installation

_Readme is based on the latest published version at the moment `v1.0.0`_

There are the following ways to install the keygen:

- Binary download from the GitHub releases page: https://github.com/stroomnetwork/keygen/releases/tag/v1.0.0
- Download as Docker image by `docker pull stroomnetwork/keygen:v1.0.0`
- (For Go users) install into GOBIN directory by `go install github.com/stroomnetwork/keygen/cmd/keygen@v1.0.0`

## How To Use

To generate a key pair and print it to console, run the following command:

```shell
keygen generate
```

To generate a key pair into a file simply invoke:

```shell
keygen generate -o keys.json
```

This will create a `keys.json` file in the workdir containing JSON with public key and private key fields.

The same command for docker image would be:

```shell
docker run -v ./:/keygen stroomnetwork/keygen:v1.0.0 generate -o tmp.json
```

_Please note that we need to correctly mount host directory with docker container._
