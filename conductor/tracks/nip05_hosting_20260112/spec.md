# Track Specification: Implement NIP-05 Hosting Server

## Overview
Implement a standalone NIP-05 hosting server in `cmd/nip05`. This server will handle HTTP requests to serve NIP-05 identifiers (Nostr addresses) mapped to public keys, allowing users to verify their identities on the Nostr network.

## Goals
- Create a new command entry point at `cmd/nip05`.
- Implement an HTTP server that responds to `/.well-known/nostr.json`.
- Support configuration via environment variables (port, data source).
- Ensure secure defaults and appropriate logging using `slog`.
- Provide a simple way to manage the mapping of names to public keys directly via environment variables.

## Detailed Requirements

### 1. Command Structure
- The application entry point must be `cmd/nip05/main.go`.
- Use `github.com/caarlos0/env` for configuration parsing.

### 2. HTTP Server
- Listen on a configurable port (default: 8080).
- Implement a handler for `GET /.well-known/nostr.json`.
- Query Parameter: `name` (the local part of the NIP-05 address).
- Response Format (JSON):
  ```json
  {
    "names": {
      "<name>": "<pubkey>"
    },
    "relays": {
      "<pubkey>": [ "wss://relay.example.com", "wss://relay2.example.com" ]
    }
  }
  ```
- CORS headers should be set appropriately to allow access from Nostr clients (`Access-Control-Allow-Origin: *`).

### 3. Configuration
- **NIP05_PORT**: Port to listen on (default: "8080").
- **NIP05_HOST**: Host interface to bind to (default: "0.0.0.0").
- **NIP05_MAPPING**: A comma-separated list of `name:pubkey` pairs to define the NIP-05 mappings (e.g., `bob:hexpubkey1,alice:hexpubkey2`).
- **LOG_LEVEL**: Logging level (debug, info, warn, error).

### 4. Data Management
- The primary data source for `name -> pubkey` mappings will be the `NIP05_MAPPING` environment variable.
- The application should parse this map at startup and store it in memory for fast lookups.
- Design the interface to be extensible for future database backends, even though the initial implementation is env-var based.

### 5. Logging
- Use `log/slog`.
- Log server startup configuration (redacting sensitive info).
- Log incoming requests and their status codes.

## Non-Functional Requirements
- **Security**: Validate the `name` parameter to prevent abuse.
- **Performance**: Efficient lookup of names (in-memory map).
- **Maintainability**: Clean separation of HTTP handling and business logic.