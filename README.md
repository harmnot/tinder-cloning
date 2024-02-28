# Before run
make sure you have installed the following tools:
- Install [Go - 1.22.0](https://www.python.org/downloads/)
- Install [Docker - 4.27.2](https://www.docker.com/products/docker-desktop)
- Install [Docker Compose - 4.27.2](https://docs.docker.com/compose/install/)
- make sure your OS support to run Makefile
- install [Dbmate](https://github.com/amacneil/dbmate) for database migration

# How to run
- Clone this repository;
- Open terminal and navigate to the project folder;
- after install `Dbmate`, run docker database by running `make run-db`, and wait until the database is up;
- Run `make migrate-db-up` to migrate database;
- Run `make run-app` to run the application;

> if you don't want to use Dbmate, you can manually the query in `migrations` folder into your database UI.
>
> also if you run the test, you can run `make test` to run the test.