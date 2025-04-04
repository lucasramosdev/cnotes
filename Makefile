# ==============================================================================
# Docker compose commands

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

init-dev:
	echo "Starting air reload..."
	docker compose up --build air