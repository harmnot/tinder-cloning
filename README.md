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

# Code Structure
The code structure is divided into several packages. Each package has its own responsibility. The following is the code structure:
- `migrations` folder contains the database migration files
- `app` package contains the main application code
- `models` package contains the models for the application
- `pkg` package contains the common packages for the application such as driver,middleware, mapper json, and util.
   - `driver` package contains the database driver / connector for database connection
   - `middleware` package contains the middleware for the application
   - `util` package contains the utility for the application
- `services` package contains the services for the application
   - `handler` package contains the handler http for the application
   - `repository` package contains the repository for the application for managing the database
   - `usecase` package contains the logic for the application
   - `schema` package contains the schema for the application, such as request and response
- `Makefile` contains the commands to run the application
- `tinder-cloning.postman_collection.json` contains the API documentation
- `Dockerfile` contains the Dockerfile for the application
- `docker-compose.yml` contains the Docker Compose file
- `go.mod` and `go.sum` contains the Go module files
- `main.go` contains the main application file
- `README.md` contains the documentation
- `.env` contains the environment variables
- `.gitignore` contains the files and folders to be ignored by Git

# API Documentation
You can use POSTMAN to test the API. The API documentation can be found on this root folder with the name `tinder-cloning.postman_collection.json`

# Database Design
![tindercloning-database-design](https://github.com/harmnot/tinder-cloning/assets/42674439/319709da-6a2c-42ba-833e-1c4050b16265)
