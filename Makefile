include .env # include .env file here

_reload:
	# reload schema cache
	docker exec -i tinderDatabase psql -d $(DB_NAME) -U $(DB_USER) -c "NOTIFY pgrst, 'reload schema'";
_migrate:
	# this migrate only run on ./$(dir) folder
    # we use 'dbmate' for migration, take a look at https://github.com/amacneil/dbmate
	# example make migrate dir=rpc tb=rpc c=up n=alter_table_a_add_col_h
	# this command use SSLMODE DISABLED !! please change it if you want to use SSLMODE
	# please change the value of DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, LOCALHOST on .env file
	# make sure .sql file has --migrate:up or --migrate:down
	# TARGET_DIRNAME condition: if directory TARGET_DIRNAME exists: take a look on conditional on the top
	# make sure you put on the top include .env
	dbmate \
		--url "postgres://$(DB_USER):$(DB_PASSWORD)@$(LOCALHOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" \
		-d "./$(dir)" --migrations-table 'schema_migrations' $(c) $(n);
migrate-db-new:
	# example: make migrate-db-new n=alter_table_a_add_col_h
	# create new file migrate db on file ./migrations on table schema_migrations_new
	$(MAKE) _migrate dir=migrations c=new n=$(n);
migrate-db-up:
	# migrate db on file ./migrations on table schema_migrations_new
	$(MAKE) _migrate dir=migrations c=up;
	$(MAKE) _reload;
migrate-db-down:
	# migrate rollback/down db on file ./migrations on table schema_migrations_new
	$(MAKE) _migrate dir=migrations c=down;
	$(MAKE) _reload;
run-db:
	# run database ui on http://localhost:5050
	docker-compose -f docker-compose.cloud.yaml up -d
    # run postgres on port 5435 ( depends on .env file )
	docker-compose up -d db
run-app:
	go run main.go
run-services:
	# run all services
	docker-compose -f docker-compose.cloud.yaml up -d
	docker-compose up -d
build-app-docker:
	# build docker image
	docker build --no-cache app
test:
	go test -count=1 -v ./...