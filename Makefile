TASK ?=

ifeq ($(strip $(TASK)),)
TARGET := main.go
else
TARGET := answers/$(TASK).go
endif

.PHONY: docker_up
docker_up:
	docker compose up -d
.PHONY: docker_stop
docker_stop:
	docker compose stop
.PHONY: run
run:
	docker compose exec -T app go run $(TARGET) < input.txt
.PHONY: fmt
fmt:
	docker compose exec app go fmt $(TARGET)
