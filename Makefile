generateJson:
	@go run generate_data/main.go

run:
	@go run .

env:
	@cp .env.example .env