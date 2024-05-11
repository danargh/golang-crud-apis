DB_URL=postgresql://postgres:180701@localhost:5432/crud_golang_apis?sslmode=disable
MIGRATION_PATH=migrations/

migration_up: 
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" -verbose up

migration_down: 
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" -verbose down

migration_fix: 
	migrate -database "$(DB_URL)" -path "$(MIGRATION_PATH)" force VERSION

show:
	go run main.go