# logger


[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4b4f93142c4d4f80ae386898a310e0e4)](https://www.codacy.com/app/me_102/logger?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Travix-International/logger&amp;utm_campaign=badger)
[![version](https://img.shields.io/github/tag/Travix-International/logger.svg)](https://github.com/Travix-International/logger) [![Build Status](https://img.shields.io/travis/Travix-International/logger/master.svg)](http://travis-ci.org/Travix-International/logger)

> Logger library in Go

## Installation

Use [`go get`](https://golang.org/cmd/go/):

```
$ go get github.com/Travix-International/logger
```

## Usage

You can import the package, and set it up with one or more transports:

```go
package main

import (
    "github.com/Travix-International/logger"
)

func main() {
    defaultMeta := make(map[string]string)
    myLogger, loggerErr := logger.New(defaultMeta)
    myLogger.AddTransport(logger.ConsoleTransport)

    // HTTP:
    // jsonFormat := logger.NewJSONFormat()
    // myLogger.AddTransport(logger.NewHttpTransport(myUrl, jsonFormat))

    // regular logs
    myLogger.Debug("EventName", "message...")
    myLogger.Info("EventName", "message...")
    myLogger.Warn("EventName", "message...")
    myLogger.Error("EventName", "message...")

    // with meta
    meta := map[string]string {
      "key": "value"
    })
    myLogger.DebugWithMeta("EventName", "message...", meta);
    myLogger.InfoWithMeta("EventName", "message...", meta);
    myLogger.WarnWithMeta("EventName", "message...", meta);
    myLogger.ErrorWithMeta("EventName", "message...", meta);

    // custom levels
    myLogger.Log("CustomLevelName", "EventName", "message...", map[string]string {
      "key": "value"
    })

    // Level filtering (define per transport!)
    filteredTransport := logger.NewTransport( ... )
    filteredTransport.SetFilter(logger.FilterByMinimumLevel(logger.NewLevelFilter("Warning")))
    myLogger.AddTransport(filteredTransport)
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

For generating coverage:

```
$ make cover
```

## License

MIT © [Travix International](https://travix.com)
