
run:
	go run cmd/main.go

gen-swag:
	swag init -g ./api/api.go -o ./api/docs

migration-up:
	migrate -path ./db/sql/ -database "postgresql://samandarxon:1234@localhost:5432/market_system?sslmode=disable" -verbose up

migration-down:
	migrate -path ./db/sql/ -database "postgresql://samandarxon:1234@localhost:5432/market_system?sslmode=disable" -verbose down
