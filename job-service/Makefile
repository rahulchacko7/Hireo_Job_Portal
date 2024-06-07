run:
	go run cmd/main.go	

mock: ##make mock files using mockgen
	mockgen -source pkg/repository/interface/interface.go -destination pkg/repository/mock/job_repository_mock.go -package mock

test:
	go test ./...