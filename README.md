# procfile-entrypoint

Launch a command based on a `Procfile` process.

## Installation

### Local

```shell
curl -L -o "procfile-entrypoint" \
  "https://github.com/acodeninja/procfile-entrypoint/releases/latest/download/procfile-entrypoint_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m)"
chmod +x procfile-entrypoint
```

## Usage

Assuming a `Procfile` with the processes `web` and `worker`:

- `procfile-entrypoint` and `procfile-entrypoint web` will start the `web` process.
- `procfile-entrypoint worker` will start the `worker` process.
