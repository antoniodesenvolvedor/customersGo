environment:
	docker-compose --env-file .env-dev up -d

run:
	go run .