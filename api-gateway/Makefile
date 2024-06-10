run:
	go run cmd/main.go


	protoc --go_out=. --go-grpc_out=. ./pkg/pb/employer/employer.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/jobseeker/jobseeker.proto


swag:
	swag init -g cmd/main.go -o ./cmd/docs