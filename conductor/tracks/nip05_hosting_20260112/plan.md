# Track Plan: Implement NIP-05 Hosting Server

## Phase 1: Setup and Configuration [checkpoint: 9464737]
- [x] Task: Initialize `cmd/nip05` directory and basic `main.go` structure. 6a2bdbe
- [x] Task: Define configuration struct with `env` tags, including the `NIP05_MAPPING` map support. 68cca62
- [x] Task: Configure `slog` based on the environment variable. fc7f50a
- [x] Task: Conductor - User Manual Verification 'Setup and Configuration' (Protocol in workflow.md) 9464737

## Phase 2: Core Logic and Data Handling [checkpoint: ada3b9a]
- [x] Task: Define an interface for NIP-05 data storage/retrieval. 52cf142
- [x] Task: Implement an in-memory data adapter that initializes from the parsed configuration map. 9faf8bc
- [x] Task: Write unit tests to verify the map parsing and lookup logic. f4a3dac
- [x] Task: Conductor - User Manual Verification 'Core Logic and Data Handling' (Protocol in workflow.md) ada3b9a

## Phase 3: HTTP Server Implementation [checkpoint: ca08ca5]
- [x] Task: Implement the HTTP handler for `/.well-known/nostr.json` according to NIP-05 spec. c8340fa
- [x] Task: Add CORS middleware handling. 7fcc7d6
- [x] Task: Write integration tests for the HTTP handler (checking correct JSON response and status codes). 69083c2
- [x] Task: Wire up the HTTP server in `main.go` with the configuration and data adapter. f762585
- [x] Task: Conductor - User Manual Verification 'HTTP Server Implementation' (Protocol in workflow.md) ca08ca5

## Phase 4: Final Polish and Documentation [checkpoint: c814304]
- [x] Task: Add graceful shutdown handling for the server. f32bfff
- [x] Task: Update project README with instructions on how to run and configure the NIP-05 server using env vars. 6656802
- [x] Task: Conductor - User Manual Verification 'Final Polish and Documentation' (Protocol in workflow.md) c814304