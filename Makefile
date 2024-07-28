include app.env

db_init:
	export PGPASSWORD=${POSTGRES_PASSWORD} && createdb -U ${POSTGRES_USER} -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} ${POSTGRES_DB}


db_migrate_up:
	migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -path db/migrations up

db_migrate_down:
	migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -path db/migrations down
