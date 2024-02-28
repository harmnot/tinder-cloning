# Introduction 
This is a simple REST API for a simple dating application like Tinder / Bumble. This application is built using Go and PostgreSQL. This application has the following features:
- User registration
- User login
- User profile
- Make reaction to other user
- Memberships (free, premium and gold) with different features

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


# API Documentation
You can use POSTMAN to test the API. The API documentation can be found on this root folder with the name `tinder-cloning.postman_collection.json`

# Database Design
![tindercloning-database-design](https://github.com/harmnot/tinder-cloning/assets/42674439/319709da-6a2c-42ba-833e-1c4050b16265)
