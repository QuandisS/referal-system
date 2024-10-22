run:
	docker compose -f deployments/docker-compose.yml --project-directory ./ up

clean-run: 
	docker compose -f deployments/docker-compose.yml --project-directory ./ up --force-recreate --renew-anon-volumes --build

.PHONY: run clean-run