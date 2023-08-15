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
  hookdeck "github.com/hookdeck/hookdeck-go-sdk"
  hookdeckclient "github.com/hookdeck/hookdeck-go-sdk/client"
)

client := hookdeckclient.NewClient(
  hookdeckclient.ClientWithAuthToken("<YOUR_AUTH_TOKEN>"),
)
```

## Usage

```go
package main

import (
  "context"
  "fmt"

  hookdeck "github.com/hookdeck/hookdeck-go-sdk"
  hookdeckclient "github.com/hookdeck/hookdeck-go-sdk/client"
)

client := hookdeckclient.NewClient(
  hookdeckclient.ClientWithAuthToken("<YOUR_API_KEY>"),
)
attempts, err := client.Attempts().GetAttempts(
  context.TODO(),
  &hookdeck.GetAttemptsRequest{
    EventId: hookdeck.String("<EVENT_ID>"),
  },
)
if err != nil {
  return err
}
fmt.Printf("Got %d attempts\n", *attempts.Count)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()
attempts, err := client.Attempts().GetAttempts(
  context.TODO(),
  &hookdeck.GetAttemptsRequest{
    EventId: hookdeck.String("<EVENT_ID>"),
  },
)
if err != nil {
  return err
}
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which includes
configuring authorization tokens to be sent on every request, or providing your own instrumented
`*http.Client`. Both of these options are shown below:

```go
client := hookdeckclient.NewClient(
  hookdeckclient.ClientWithAuthToken("<YOUR_AUTH_TOKEN>"),
  hookdeckclient.ClientWithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
attempts, err := client.Attempts().GetAttempts(
  context.TODO(),
  &hookdeck.GetAttemptsRequest{
    EventId: hookdeck.String("<EVENT_ID>"),
  },
)
if err != nil {
  if badRequestErr, ok := err.(*hookdeck.BadRequestError);
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
attempts, err := client.Attempts().GetAttempts(
  context.TODO(),
  &hookdeck.GetAttemptsRequest{
    EventId: hookdeck.String("<EVENT_ID>"),
  },
)
if err != nil {
  var badRequestErr *hookdeck.BadRequestError
  if errors.As(err, badRequestErr); ok {
    // Do something with the bad request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
attempts, err := client.Attempts().GetAttempts(
  context.TODO(),
  &hookdeck.GetAttemptsRequest{
    EventId: hookdeck.String("<EVENT_ID>"),
  },
)
if err != nil {
  return fmt.Errorf("failed to get attempts: %w", err)
}
```

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
