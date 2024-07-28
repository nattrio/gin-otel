.PHONY: up down down-v app
up:
	cd docker && docker compose up --build -d
down:
	cd docker && docker compose down
down-v:
	cd docker && docker compose down -v
app:
	UPTRACE_DSN=http://project2_secret_token@localhost:14317/2 go run .
	

.PHONY: pg-create pg-up pg-down
pg-create:
	@read -p "Enter postgres migration name: " name; \
	migrate create -ext sql -dir db/migration $$name
pg-up:
	bash scripts/pg-up
pg-down:
	bash scripts/pg-down

.PHONY: check-swagger swagger swagger-serve
check-swagger:
	which swagger
swagger: check-swagger
	swagger generate spec -o ./docs/swagger.yaml --scan-models
swagger-serve: check-swagger swagger
	swagger serve -F=swagger ./docs/swagger.yaml --no-open

.PHONY: dep-tree
dep-tree:
	mkdir -p .dep-tree
	@read -p "Enter path: " path; \
	dep-tree entropy --no-browser-open --enable-gui --render-path ./.dep-tree/index.html $$path
