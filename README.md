# lc-integration-testing

Integration tests for limpidchart services.

To run tests locally you must start services. You can start develop or stable release using docker compose.

Do the following to start the develop release:

```sh
docker compose --file compose/develop/docker-compose.yml pull
docker compose --file compose/develop/docker-compose.yml up -d
```

Do the following to start the stable release:

```sh
docker compose --file compose/stable/docker-compose.yml pull
docker compose --file compose/stable/docker-compose.yml up -d
```

Then you can run the tests:

```sh
go test -v ./internal/integration/...
```

Don't forget to clean up after running tests:

```sh
docker compose --file compose/develop/docker-compose.yml down --remove-orphans
```

```sh
docker compose --file compose/stable/docker-compose.yml down --remove-orphans
```
