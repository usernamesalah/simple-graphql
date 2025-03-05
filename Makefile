.PHONY: dev
dev:
	@echo "> Run Weather Forecast API for Development with default config ..."
	@docker compose --project-directory ./deployments/development -p tensor-graphql $(args)

.PHONY: dev-migrate
dev-migrate:
	@echo "> Running database migration ..."
	@docker exec tensor-graphql-development ./docker/development/db-migration.sh $(args)

mock:
	@./scripts/generate_mocks.sh

mock-win:
	@powershell scripts/generate_mocks.sh

test-report: 
	go test ./internal/... -v -coverprofile cover.out
	go tool cover -html=cover.out

test:
	go test ./internal/... -v

gen-swagger:
	@echo "Updating API documentation..."
	@swag init -o ${API_DOCS_PATH} -g cmd/webservice/main.go