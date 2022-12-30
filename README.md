# shared

Shared services or configs


# Notes

To access Portainer without Traefik port forward `9443` to your machine.

For accass, the credentials are available in `.env` file

## etcd

1. Build the server: `make etcd`
2. Run the server: `./etcd`

The HTTP server is available on `localhost:8080`:
- `/configs` - Returns the list of available configs
- `/config/{id}` - Returns the config with the given id (key)

The server is also accesible via gRPC on `localhost:9090`.

To add, remove or update a config, simple add, delete or edit a file
in the `configs` directory with the value(s) you want.