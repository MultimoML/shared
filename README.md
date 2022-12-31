# shared

Shared services or configs

# Notes

To access Portainer without Traefik port forward `9443` to your machine.

For access, the credentials are available in `.env` file

## etcd

1. Build the server: `make etcd`
2. Run the server: `./etcd` (run it in the same directory as the `configs` directory)

The HTTP server is available on `https://configs.multimo.ml`:
- `/live` - Whether the server is alive
- `/ready` - Whether the server is ready
- `/` - Returns the list of available configs
- `/{id}` - Returns a config with the given id/key if found

The server is also accessible via gRPC on `https://configs.multimo.ml:9090`.

Configs are stored in the `configs` directory (create it if it doesn't exist yet),
next to the binary.
To add, remove or update a config, simple add, delete or edit a file
in the `configs` directory with the value(s) you want.