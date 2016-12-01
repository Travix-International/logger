# logger

[![version](https://img.shields.io/github/tag/Travix-International/logger.svg)](https://github.com/Travix-International/logger) [![Build Status](https://img.shields.io/travis/Travix-International/logger/master.svg)](http://travis-ci.org/Travix-International/logger)

> Logger library in Go

## Installation

Use [`go get`](https://golang.org/cmd/go/):

```
$ go get github.com/Travix-International/logger
```

Or, [`gvt`](https://github.com/FiloSottile/gvt):

```
$ gvt fetch github.com/Travix-International/logger
```

## Usage

You can import the package, and set it up with one or more transports:

```go
package main

import (
    "github.com/Travix-International/logger"
    "github.com/Travix-International/logger/transports/console"
)

func main() {
    myLogger := logger.New()
    myLogger.AddTransport(console.New())

    // regular logs
    myLogger.Debug("EventName", "message...")
    myLogger.Info("EventName", "message...")
    myLogger.Warn("EventName", "message...")
    myLogger.Error("EventName", "message...")

    // error objects
    myLogger.Exception("EventName", err, "message...")

    // with meta
    meta := myLogger.Meta()
    meta.Set("key", "value")
    myLogger.
        WithMeta(meta).
        Info("EventName", "message...")
}
```

## Development

Clone the repo to your `$GOPATH`:

```
$ git clone git@github.com:Travix-International/logger.git $GOPATH/src/github.com/Travix-International/logger
```

Run tests:

```
$ make run-tests
```

## License

MIT © [Travix International](https://travix.com)
