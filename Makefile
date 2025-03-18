.PHONY: docker_up
docker_up:
	docker compose up -d
.PHONY: docker_down
docker_down:
	docker compose down
.PHONY: run
run:
	docker compose exec -T app go run main.go < input.txt
.PHONY: fmt
fmt:
	docker compose exec app go fmt main.go