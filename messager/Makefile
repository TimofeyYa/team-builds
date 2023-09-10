get_env:
	source ./.env

#make migrate_up DB_USER=user DB_PASS=pass DB_HOST=localhost DB_PORT=5433 DB_NAME=messages
migrate_up:
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" status
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up
migrate_down:
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" status
	goose -dir ./migrations postgres "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

run:
	go run ./cmd/main.go