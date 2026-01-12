# nostr-tools

A suite of tools to easily and securely build and operate external services such as Nostr relays and NIP-05.

## Components

### NIP-05 Hosting Server (`cmd/nip05`)

A standalone server that handles HTTP requests to serve NIP-05 identifiers.

#### Configuration

The server is configured entirely via environment variables:

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `NIP05_PORT` | Port to listen on | `8080` | No |
| `NIP05_HOST` | Host interface to bind to | `0.0.0.0` | No |
| `NIP05_MAPPING` | Comma-separated `name:pubkey` pairs | | Yes |
| `LOG_LEVEL` | Logging level (`debug`, `info`, `warn`, `error`) | `info` | No |

Example `NIP05_MAPPING`: `bob:73c91d8...,alice:83d12...`

#### Running

```bash
# Build the binary
go build -o nip05 ./cmd/nip05

# Run with environment variables
export NIP05_MAPPING="yourname:yourhexpubkey"
./nip05
```

#### NIP-05 Verification Endpoint

The server serves the standard NIP-05 endpoint:
`GET /.well-known/nostr.json?name=<name>`

Example:
`curl "http://localhost:8080/.well-known/nostr.json?name=yourname"`

## Development

### Requirements
- Go 1.25.5 or later

### Running Tests
```bash
go test ./...
```

### Linting
```bash
golangci-lint run
```
