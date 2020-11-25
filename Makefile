build: ## Build the application
	CGO_ENABLED=0 go build -o dist/server go-http-server/server.go

run: ## Run the application
	go run go-http-server/server.go