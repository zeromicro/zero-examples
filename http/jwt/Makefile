build:api
	@go build -o cmd

run:
	@go run jwt.go -f etc/jwt-api.yaml

api:clean_api
	@goctl api go -api jwt.api -dir .

clean_api:
	@echo "doing clean_api"
	@rm -rf ./internal/handler/*

clean_model:
	@rm -rf ./model/*