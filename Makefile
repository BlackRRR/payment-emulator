run:
	docker-compose up -d --force-recreate --build

tests:
	go test ./test/service_transaction_test.go
