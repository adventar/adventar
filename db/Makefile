start:
	docker-compose --project-name adventar-db up

mysql:
	mysql -u root -h 127.0.0.1 -P 13306 adventar_dev

migrate-dry-run:
	mysqldef -u root -h 127.0.0.1 -P 13306 adventar_dev --dry-run < schema.sql

migrate-dev:
	mysqldef -u root -h 127.0.0.1 -P 13306 adventar_dev < schema.sql

migrate-test:
	mysql -u root -h 127.0.0.1 -P 13306 -e 'CREATE DATABASE IF NOT EXISTS `adventar_test`'
	mysqldef -u root -h 127.0.0.1 -P 13306 adventar_test < schema.sql

migrate: migrate-dev migrate-test
