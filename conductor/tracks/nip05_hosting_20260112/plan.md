# Track Plan: Implement NIP-05 Hosting Server

## Phase 1: Setup and Configuration
- [x] Task: Initialize `cmd/nip05` directory and basic `main.go` structure. 6a2bdbe
- [x] Task: Define configuration struct with `env` tags, including the `NIP05_MAPPING` map support. 68cca62
- [ ] Task: Configure `slog` based on the environment variable.
- [ ] Task: Conductor - User Manual Verification 'Setup and Configuration' (Protocol in workflow.md)

## Phase 2: Core Logic and Data Handling
- [ ] Task: Define an interface for NIP-05 data storage/retrieval.
- [ ] Task: Implement an in-memory data adapter that initializes from the parsed configuration map.
- [ ] Task: Write unit tests to verify the map parsing and lookup logic.
- [ ] Task: Conductor - User Manual Verification 'Core Logic and Data Handling' (Protocol in workflow.md)

## Phase 3: HTTP Server Implementation
- [ ] Task: Implement the HTTP handler for `/.well-known/nostr.json` according to NIP-05 spec.
- [ ] Task: Add CORS middleware handling.
- [ ] Task: Write integration tests for the HTTP handler (checking correct JSON response and status codes).
- [ ] Task: Wire up the HTTP server in `main.go` with the configuration and data adapter.
- [ ] Task: Conductor - User Manual Verification 'HTTP Server Implementation' (Protocol in workflow.md)

## Phase 4: Final Polish and Documentation
- [ ] Task: Add graceful shutdown handling for the server.
- [ ] Task: Update project README with instructions on how to run and configure the NIP-05 server using env vars.
- [ ] Task: Conductor - User Manual Verification 'Final Polish and Documentation' (Protocol in workflow.md)