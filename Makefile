TARGET ?=

ifeq ($(strip $(TARGET)),)
GO_FILE := main.go
else
GO_FILE := answers/$(TARGET).go
endif

.PHONY: docker_up
docker_up:
	docker compose up -d
.PHONY: docker_stop
docker_stop:
	docker compose stop
.PHONY: run
run:
	docker compose exec -T app go run $(GO_FILE) < input.txt
.PHONY: fmt
fmt:
	docker compose exec app go fmt $(GO_FILE)
