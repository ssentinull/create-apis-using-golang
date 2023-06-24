# command to run the server in daemon mode
run-server:
	@modd -f ./.modd/server.modd.conf

# command to run migration
# eg: make migrate direction=up
migrate-db:
	go run internal/cmd/migration/main.go -direction=$(direction) -step=0

# command to run db seeder based on number of $(seed)
# eg: make seed-db seed=10
seed-db:
	go run internal/cmd/seeder/main.go -seed=$(seed)
