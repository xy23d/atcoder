.PHONY: docker_up
docker_up:
	docker compose up -d
.PHONY: docker_stop
docker_stop:
	docker compose stop
.PHONY: run
run:
	docker compose exec -T app go run main.go < input.txt
.PHONY: fmt
fmt:
	docker compose exec app go fmt main.go
