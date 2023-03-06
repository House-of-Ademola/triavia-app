
# Open connection to the backend
backend-up:
	docker compose up

# Close connection to the backend
backend-down:
	docker compose down

# Run each time there are changes made to the Dockerfile / and Compose
docker-build:
	docker compose build

.PHONY: docker-up docker-down