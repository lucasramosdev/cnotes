# ==============================================================================
# Environment commands

set-env:
	export CNOTES_DBUSER="user" && export CNOTES_DBPASS="root" && export CNOTES_DB="database"

# ==============================================================================
# Docker compose commands

clean-db:
	echo "Cleaning db data..."
	docker volume rm cnotes_dbdata

up-db: set-env
	echo "Starting docker compose with only db app..."
	docker compose up -d db