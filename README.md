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

# Sequence Diagram
![tinder-cloning-sequence](https://github.com/harmnot/tinder-cloning/assets/42674439/0ac8e3ff-0325-4754-93ae-00093d3a9920) \
[Sequence Codes by PlantUML](https://www.plantuml.com/plantuml/uml/xLRDRjim3BxhANZhTYYsNNleEdJOBZqCsnCm5XqBbIMFb9FcxMU5jXCLiRVULg06uWF5x-CFFp6-auIS-jQRBSU-WtiKWgkVj7MIYA_ClNSKqlKjVEcw6_Zcl8SwSYi6VO8TUcSnaPh0Oa93NWXSE1wf9KElRvNe1gI9Uo3coD2I26v959DzbN54V1qC6nIwWIqbw8LOVJO7e1QTXmz7Oy2M30nFpgXgg7fKN_HCQ0VV9oVfim5z9jBhssD1J1Mv7K5gICbThG28324UKLhHQ1zvBjIarMGgJ6K6F7QJRC4NOYeubLJpQb0Qw51a7zcUpiks2Ewch4EAOCS30nTpxAkAHEPaD9LWpVGh1F0GjouiJn_UZIF0OESquVr_80gpC66hRi_ZKh1d93k7aE8ZXSjD-Ku9Vd2idITvtfjyzk0awIW6apd9qTM3jVAzfP4Tm9n4CDcM5SNqktTCLZqLsg1CmFD0QfiwDWMPEg-0LVcMlafHaXOwtH6x-Z0wP37GKsAi4gHcX6b3vXwR2pe-kxGjjHjYQLntGTCBaor1TRmO8lOXGTKpKr3PP7c2x5r7aiMJeiL3HgxhHtaKTbewiUsGwIoFufxNJAGXCTY2OTuGdpwQlwsLcMK6ASU0XUG6BotXNoE7c3Aog1mkN34j6oi5xnjNDHchD1OsL4UctCk2Pwc7MHK-tv9Ehz5IrOMJ53LrJAthMuWCs-ZFX_yAQWWMpBnqo-NoAoUpk5xqD8TEo5kQUT-sp6lVCczRPbPLAUrd6PGawzTpbe1spidBgAV2Ig_2PQs7Qx7iKgpSgGEbrK-zgJIbcbtBJkWt3PDqCGYjf3jz44vfZwuP_OA5-qgc-JRDygydl-w6Lz_wFm00)
