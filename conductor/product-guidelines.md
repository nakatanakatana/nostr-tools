# Product Guidelines

## Tone and Style
- **Minimalist**: Documentation and messages should be concise and to the point, avoiding fluff.

## Logging and Output
- **Structured Logging**: Use `slog` for structured logging.
- **Log Levels**: Assign appropriate log levels (Debug, Info, Warn, Error) to log entries.
- **Minimal Default**: By default, output only the minimum necessary logs (e.g., Info or Warn level and above).

## Design Principles
- **Secure by Default**: Security features must be enabled by default. Secure operation should be the baseline without requiring extra configuration.
- **Environment-Based Configuration**: All configuration options must be acceptable via environment variables to support containerized deployments and 12-Factor App principles.
