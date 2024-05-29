APP_NAME := thn-rest-api
PKG := ./...
GREEN := \033[0;32m
NC := \033[0m # No Color

.PHONY: all deps test build run clean
all: deps test build

deps:
	@echo -e "$(GREEN)Instalando dependencias...$(NC)"
	@go mod tidy

test:
	@echo -e "$(GREEN)Ejecutando pruebas...$(NC)"
	@go test $(PKG) -v

build:
	@echo -e "$(GREEN)Compilando el proyecto...$(NC)"
	@go build -o $(APP_NAME) ./cmd/api/main.go

run: build
	@echo -e "$(GREEN)Ejecutando la aplicación...$(NC)"
	@./$(APP_NAME)

clean:
	@echo -e "$(GREEN)Limpieza de archivos generados...$(NC)"
	@rm -f $(APP_NAME)

test-unit:
	@echo -e "$(GREEN)Ejecutando pruebas unitarias...$(NC)"
	@go test $(PKG) -run Unit -v

test-integration:
	@echo -e "$(GREEN)Ejecutando pruebas de integración...$(NC)"
	@go test $(PKG) -run Integration -v