# 默认目标
.PHONY: build clean test help

# 应用名称
APP_NAME := ca-import-tool

# Go参数
GO_BUILD := go build
GO_CLEAN := go clean
GO_TEST := go test

# 默认构建
build: build-windows build-darwin build-linux

# 构建Windows版本
build-windows:
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-windows.exe

# 构建macOS版本
build-darwin:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-darwin

# 构建Linux版本
build-linux:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-linux

# 清理构建文件
clean:
	$(GO_CLEAN)
	rm -rf bin/

# 运行测试
test:
	$(GO_TEST) -v ./...

# 本地构建
build-local:
	$(GO_BUILD) -o $(APP_NAME)

# 显示帮助信息
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@echo "  build            构建所有平台版本"
	@echo "  build-windows    构建Windows版本"
	@echo "  build-darwin     构建macOS版本"
	@echo "  build-linux      构建Linux版本"
	@echo "  build-local      本地构建"
	@echo "  clean            清理构建文件"
	@echo "  test             运行测试"
	@echo "  help             显示此帮助信息"