.DEFAULT_GOAL := help
.PHONY: checkout-master

checkout-master:
		git fetch && git checkout master && git pull origin master

migrate-up:
		migrate -path=internal/postgres/migrations \
				-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
				-verbose up

migrate-down:
		migrate -path=internal/postgres/migrations \
				-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
				-verbose down
