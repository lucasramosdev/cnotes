# ==============================================================================
# Docker compose commands

down-containers:
	echo "Removing containers..."
	docker compose down

clean-db:
	echo "Cleaning db data..."
	docker volume rm cnotes_dbdata

up-db:
	echo "Starting docker compose with only db app..."
	docker compose up -d db

migrate-db:
	echo "Migrating db app..."
	docker compose up -d db migrate

up-app:
	echo "Starting docker compose app..."
	docker compose up --build -d app db migrate

up-dev:
	echo "Starting air reload..."
	docker compose up air

build-dev:
	echo "Starting air reload..."
	docker compose up --build air