PATH_TO_ENV_FILE := "/Users/user/go/src/space_service/.env"

# NASA GRPC SERVICE
SERVICE_ENTRY_POINT := nasa_grpc_service/cmd/main.go
SERVICE_COVERPKG 	:= ./nasa_grpc_service/...
SERVICE_OUT_FILE 	:= ./grpc_cov_report/nasa_grpc_service.out
SERVICE_TEST_DIR 	:= ./nasa_grpc_service/test
SERVICE_HTML_FILE 	:= ./grpc_cov_report/nasa_grpc_service.html

# SPACE HTTP WEB APP
HTTP_ENTRY_POINT := space_web_app/cmd/main.go
HTTP_COVERPKG 	 := ./space_web_app/...
HTTP_OUT_FILE 	 := ./http_cov_report/space_web_app.out
HTTP_TEST_DIR 	 := ./space_web_app/test
HTTP_HTML_FILE 	 := ./http_cov_report/space_web_app.html


run-nasa-grpc-service:
	go run $(SERVICE_ENTRY_POINT)
test-nasa-grpc-service:
	go test -coverpkg=$(SERVICE_COVERPKG) -coverprofile=$(SERVICE_OUT_FILE) $(SERVICE_TEST_DIR) && go tool cover -func=$(SERVICE_OUT_FILE) && go tool cover -html=$(SERVICE_OUT_FILE)

run-space-web-app:
	go run $(HTTP_ENTRY_POINT)
test-space-web-app:
	go test -coverpkg=$(HTTP_COVERPKG) -coverprofile=$(HTTP_OUT_FILE) $(HTTP_TEST_DIR) && go tool cover -func=$(HTTP_OUT_FILE) && go tool cover -html=$(HTTP_OUT_FILE)