# lc-intergration-testing

This repo contains integration tests for limpidchart services.

To run tests locally you must start services. You can start develop or stable release using docker compose.

Do the following to start the develop release:

```sh
cd compose/develop
docker compose up -d
cd ../..
```

Do the following to start the stable release:

```sh
cd compose/stable
docker compose up -d
cd ../..
```

Then you can run the tests:

```sh
go test -v ./internal/intergration/...
```

Don't forget to clean up after running tests:

```sh
cd compose/develop
docker compose down --remove-orphans
cd ../..
```

```sh
cd compose/stable
docker compose down --remove-orphans
cd ../..
```
