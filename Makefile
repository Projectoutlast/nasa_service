# NASA GRPC SERVICE
SERVICE_COVERPKG := ./nasa_grpc_service/...
SERVICE_OUT_FILE := nasa_grpc_service.out
SERVICE_TEST_DIR := ./nasa_grpc_service/test
SERVICE_HTML_FILE := nasa_grpc_service.html

run-nasa-grpc-service:
	go run nasa_grpc_service/cmd/main.go
test-nasa-grpc-service:
	go test -coverpkg=$(SERVICE_COVERPKG) -coverprofile=$(SERVICE_OUT_FILE) $(SERVICE_TEST_DIR) && go tool cover -func=$(SERVICE_OUT_FILE) && go tool cover -html=$(SERVICE_OUT_FILE) -o $(SERVICE_HTML_FILE)