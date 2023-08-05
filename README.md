# Fastly CLI

```shell
make run ARGS="-h"
```

## TODO

- Figure out strategies and patterns.
  - API mapping commands:
    - Validate flags (e.g. can't use --verbose and --json at same time)
    - Print output (normal, verbose, JSON)
  - Error handling (e.g. remediation errors design).
  - Compute command for setting up and creating a service.
- Possible generics (or more likely code-generation from OpenAPI schemas) for API mapping commands.
  - Otherwise manually implement each API mapping command 🙈
