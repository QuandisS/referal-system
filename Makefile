run:
	docker compose -f deployments/docker-compose.yml --project-directory ./ up

clean-run: 
	docker compose -f deployments/docker-compose.yml --project-directory ./ up --force-recreate --renew-anon-volumes --build

stop-all:
	docker compose -f deployments/docker-compose.yml --project-directory ./ down

run-db: 
	docker compose -f deployments/docker-compose.yml --project-directory ./ up -d postgres

.PHONY: run clean-run stop-all run-db