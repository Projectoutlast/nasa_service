run-nasa-grpc-service:
	go run nasa_grpc_service/cmd/main.go
run-test:
	go test -cover -race nasa_grpc_service/test