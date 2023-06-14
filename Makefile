DB_URL=postgres://match:match@db:5432/dev?sslmode=disable

migrateup:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path=sql/migrations -database "$(DB_URL)" -verbose drop

test:
	go test -cover ./...

test-clean:
	go clean --testcache

.PHONY:  migrateup migratedown test test-clean
