include .env

migrateup:
	migrate -path pkg/repository/migrations -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/repository/migrations -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

createmigrate:
	migrate create -ext sql -dir pkg/repository/migrations -seq ${file}

service:
	touch pkg/api/${file}.go
	touch pkg/app/handler${file}.go
	touch pkg/repository/repo${file}.go