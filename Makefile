# command to run the server in daemon mode
.PHONY: run-server
run-server:
	@modd -f ./.modd/server.modd.conf

# command to run migration
# eg: make migrate direction=up
.PHONY: migrate-db
migrate-db:
	go run internal/cmd/migration/main.go -direction=$(direction) -step=0

# command to run db seeder based on number of $(seed)
# eg: make seed-db seed=10
.PHONY: seed-db
seed-db:
	go run internal/cmd/seeder/main.go -seed=$(seed)

# command to generate mock interfaces
.PHONY: mockgen
mockgen:
	@command -v "mockgen" >/dev/null 2>&1 || go install github.com/golang/mock/mockgen@v1.6.0
	@rm -rf internal/model/mock
	@mockgen -destination=internal/model/mock/book.go -package=mock -source=internal/model/book.go BookRepository
	@mockgen -destination=internal/model/mock/cache.go -package=mock -source=internal/model/cache.go CacheRepository

# command to run unit tests
.PHONY: test
test:
	mkdir -p test-reports
	go test -coverprofile=test-reports/coverage.out ./internal/...
	go tool cover -func test-reports/coverage.out -o test-reports/coverage.cov

# command to generate unit test coverage html
.PHONY: test-coverage
test-coverage:
	@$(MAKE) test
	go tool cover -html=test-reports/coverage.out -o test-reports/coverage.html