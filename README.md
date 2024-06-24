# go-cloudbuild-template
Showcase a Go GraphQL API repository structure

## Setup local environment

### Requirements
 - make
 - go version 1.18
 - docker

### Setup a local environment

You can create a local mysql 8 database instance and a redis instance using:
```
make setup-local-env
make migrate-up
```

And then you can start and stop your database:
```
make start-local-env
make stop-local-env
```

### Run database migration

Database migrations are located in the `sql/migrations` directory.
We use the *golang-migrate* tool to handle migrations: https://github.com/golang-migrate/migrate

You can run migration **up** and **down** one step using:
```
make migration-up
make migration-down
```

### Local docker-compose

You can run the following command from the root of the repository
to create the whole environment in a docker container running the following command:
```
make run
```
The API will be running locally on the port :80

### TODO:

- Add Postman files
- make sure it's running
- Add basic feature like notes/auth
- Unit test
- Add CI
- POC Prisma