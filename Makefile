up:
	docker compose up --build -d

down:
	docker compose down

log-backend:
	docker compose logs -f backend

log-postgres:
	docker compose logs -f postgres

log-redis:
	docker compose logs -f redis

log-migrate:
	docker compose logs -f migrate