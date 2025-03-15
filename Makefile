.PHONY: docker_up
docker_up:
	docker compose up -d
.PHONY: docker_down
docker_down:
	docker compose down
.PHONY: docker_exec
docker_exec:
	docker compose exec app /bin/bash