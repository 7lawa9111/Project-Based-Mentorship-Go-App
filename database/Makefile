postgres:
	docker run --name postgresSQL -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=0000 -d postgres
createdb:
 	docker exec -it postgresSQL createdb --username=root --owner=root document_system
dropdb:
 	docker exec -it postgresSQL dropdb
migrateup:
	migrate -path database/migration -database "postgresql://root:0000@localhost:5432/document_system?sslmode=disable" -verbose up
migratedown:
    migrate -path database/migration -database "postgresql://root:0000@localhost:5432/document_system?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrateup migratedown
