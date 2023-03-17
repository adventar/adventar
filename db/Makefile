start:
	docker-compose --project-name adventar-db up

mysql:
	mysql -u root -h 127.0.0.1 -P 13306 adventar_dev

migrate-dry-run:
	mysqldef -u root -h 127.0.0.1 -P 13306 adventar_dev --dry-run < schema.sql

migrate:
	mysqldef -u root -h 127.0.0.1 -P 13306 adventar_dev < schema.sql
