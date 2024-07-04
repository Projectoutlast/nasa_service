PATH_TO_ENV_FILE := "/Users/user/go/src/space_service/.env"

# AUTH SERVICE
AUTH_ENTRY_POINT := auth_service/cmd/auth/main.go
AUTH_COVERPKG 	:= ./auth_service/...
AUTH_OUT_FILE 	:= ./auth_cov_report/auth_grpc_service.out
AUTH_TEST_DIR 	:= ./auth_service/test
AUTH_HTML_FILE 	:= ./auth_cov_report/auth_grpc_service.html

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
run-space-web-app:
	go run $(HTTP_ENTRY_POINT)
run-auth-service:
	go run $(AUTH_ENTRY_POINT)

test-nasa-grpc-service:
	go test -coverpkg=$(SERVICE_COVERPKG) -coverprofile=$(SERVICE_OUT_FILE) $(SERVICE_TEST_DIR) && go tool cover -func=$(SERVICE_OUT_FILE) && go tool cover -html=$(SERVICE_OUT_FILE) -o $(SERVICE_HTML_FILE)
test-space-web-app:
	go test -coverpkg=$(HTTP_COVERPKG) -coverprofile=$(HTTP_OUT_FILE) $(HTTP_TEST_DIR) && go tool cover -func=$(HTTP_OUT_FILE) && go tool cover -html=$(HTTP_OUT_FILE) -o $(HTTP_HTML_FILE)
test-auth-service:
	go test -coverpkg=$(AUTH_COVERPKG) -coverprofile=$(AUTH_OUT_FILE) $(AUTH_TEST_DIR) && go tool cover -func=$(AUTH_OUT_FILE) && go tool cover -html=$(AUTH_OUT_FILE) -o $(AUTH_HTML_FILE)

migrate-auth:
	go run ./auth_service/cmd/migrator --storage-path=./auth_service/storage/auth.db --migrations-path=./auth_service/migrations

# DOCKER RUNNER
run-app:
	docker-compose up -d

run-all-services:
	./start.sh