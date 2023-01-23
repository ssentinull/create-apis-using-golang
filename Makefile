# command to run the server in daemon mode
run-server:
	@modd -f ./.modd/server.modd.conf

# command to run migration up
migrate-up:
	go run internal/cmd/migration/main.go -direction=up -step=0

# command to run migration down
migrate-down:
	go run internal/cmd/migration/main.go -direction=down -step=0