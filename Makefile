default:

start:
	@echo "=============starting MQ Service============="
	sh scripts/start-mq-service.sh
stop:
	sh scripts/stop-mq-service.sh
start-receiver:
	go run cmd/roundrobin_receiver/main.go
start-sender:
	go run cmd/roundrobin_sender/main.go
start-broadcast-receiver:
	go run cmd/broadcast_receiver/main.go
start-broadcast-sender:
	go run cmd/broadcast_sender/main.go