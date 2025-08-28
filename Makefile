export MYSQL_URL='mysql://root:passwordje@tcp(localhost:3306)/golang'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq ${name}

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down

migrate-reset:
	@ migrate -path scripts/migrations -database ${MYSQL_URL} force 0
