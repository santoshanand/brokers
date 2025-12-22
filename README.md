# Brokers

A collection of small Go client wrappers for Indian broker APIs. Each folder contains a package implementing API helpers and models for a specific broker.

Supported providers
- `angelone` — Angel One client and models (files: `angelone.go`, `model.go`).
- `dhan` — Dhan client, orders, portfolio and market data helpers (files include `dhan.go`, `order.go`, `portfolio.go`, `marketfeed.go`, `historical_data.go`, `fund.go`, `super_order.go`, `option_chain.go`, `model.go`, `const.go`).
- `fyers` — Fyers client and enums (files: `fyers.go`, `enum.go`, `model.go`, `const.go`).
- `paytm` — Paytm client helpers (`paytm.go`, `model.go`).
- `upstox` — Upstox client (`upstox.go`, `model.go`).
- `zerodha` — Zerodha client (`zerodha.go`, `model.go`).

Getting started

Requirements:
- Go 1.18+ (module-enabled environment).

Build:

```bash
go build ./...
```

Run tests:

```bash
go test ./...
```

Usage

Each provider is a standalone package. Import the package you need and use the exposed client functions and models. Example:

```go
import (
    "github.com/santoshanand/brokers/angelone"
)

// create and use client according to package docs
```

Notes

- This repo contains small adapters and helper models for interacting with broker APIs — check individual packages for more detailed usage and examples.
- No license file included. Add `LICENSE` if you want to publish/reuse this code.

Contributing

- Open an issue or PR with suggested changes. Keep commits focused and add tests where applicable.

Contact

- Maintainer: repository owner
