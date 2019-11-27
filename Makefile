default:

start:
	@echo "=============starting MQ Service============="
	sh scripts/start-mq-service.sh
stop:
	sh scripts/stop-mq-service.sh
start-receiver:
	go run cmd/receiver/main.go
start-sender:
	go run cmd/sender/main.go