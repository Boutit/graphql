run.local:
	ENV=local go run cmd/server/main.go

run.development:
	ENV=development go run cmd/server/main.go

run.staging:
	ENV=staging go run cmd/server/main.go

run.production:
	ENV=production go run cmd/server/main.go