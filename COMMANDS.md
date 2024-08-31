### Running PostgreSQL Commands in Docker

To interact with your PostgreSQL database running inside a Docker container, you can use the following `docker compose`
command:

```sh
docker compose exec -it db psql -U baloo -d lenslocked
```

This command does the following:

- `docker compose exec -it db`: Executes commands in the `db` service container with interactive terminal.
- `psql -U baloo -d lenslocked`: Connects to the `lenslocked` database using `psql` as the `baloo` user.



