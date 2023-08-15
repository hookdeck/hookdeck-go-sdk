# HookDeck Go Library

The HookDeck Go library provides convenient access to the HookDeck API from Go.

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/hookdeck/hookdeck-go-sdk)

## Requirements

This module requires Go version >= 1.19.

## Installation

Run the following command to use the HookDeck Go library in your Go module:

```sh
go get github.com/hookdeck/hookdeck-go-sdk
```

This module requires Go version >= 1.19.

## Instantiation

```go
import (
  "context"
  "fmt"

  hookdeck "github.com/hookdeck/hookdeck-go-sdk"
  hookdeckclient "github.com/hookdeck/hookdeck-go-sdk/client"
)

client := hookdeckclient.NewClient(
  hookdeckclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  hookdeckclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
```

## Usage

```go
package main

import (
  "context"
  "fmt"
  "net/http"

  hookdeck "github.com/hookdeck/hookdeck-go-sdk"
  hookdeckClient "github.com/hookdeck/hookdeck-go-sdk/client"
)

func run() error {
  client := hookdeckClient.NewClient(
    hookdeckClient.ClientWithAuthToken("<YOUR_API_KEY>"),
  )
  attempts, err := client.Attempts().GetAttempts(
    context.TODO(),
    &hookdeck.GetAttemptsRequest{
      EventId: hookdeck.String("<event_id>"),
    },
  )
  if err != nil {
    return err
  }
}
```

## Timeouts

TODO

## Client Options

TODO

## Errors

TODO

## Beta status

This SDK is in beta, and there may be breaking changes between versions without a major version update.
Therefore, we recommend pinning the package version to a specific version in your `go.mod` file. This way,
you can install the same version each time without breaking changes unless you are intentionally looking
for the latest version.

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions
made directly to this library would have to be moved over to our generation code, otherwise they would be
overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we
will not be able to merge it as-is. We suggest opening an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
