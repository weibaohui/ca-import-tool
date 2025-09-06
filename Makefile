# 默认目标
.PHONY: build clean test help

# 应用名称
APP_NAME := ca-import-tool

# Go参数
GO_BUILD := go build
GO_CLEAN := go clean
GO_TEST := go test

# 默认构建 - 清理并构建所有平台和架构版本
build: clean build-windows build-windows-arm64 build-darwin build-darwin-arm64 build-linux build-linux-arm64 build-linux-arm

# 构建Windows版本 (AMD64)
build-windows:
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-windows-amd64.exe

# 构建Windows版本 (ARM64)
build-windows-arm64:
	GOOS=windows GOARCH=arm64 $(GO_BUILD) -o bin/$(APP_NAME)-windows-arm64.exe

# 构建macOS版本 (AMD64)
build-darwin:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-darwin-amd64

# 构建macOS版本 (ARM64)
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o bin/$(APP_NAME)-darwin-arm64

# 构建Linux版本 (AMD64)
build-linux:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o bin/$(APP_NAME)-linux-amd64

# 构建Linux版本 (ARM64)
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o bin/$(APP_NAME)-linux-arm64

# 构建Linux版本 (ARM)
build-linux-arm:
	GOOS=linux GOARCH=arm $(GO_BUILD) -o bin/$(APP_NAME)-linux-arm

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

# 生成测试证书
test-cert:
	cd test && ./generate-cert.sh

# 构建测试环境Docker镜像
test-docker-build:
	docker build -t ca-test-server test/
	@echo "Docker镜像构建完成。使用 'make test-docker-run' 运行测试环境"

# 运行测试环境Docker容器
test-docker-run:
	docker run -d -p 80:80 -p 443:443 --name ca-test ca-test-server
	@echo "测试环境已启动。访问 https://test.example.com 测试证书信任状态"

# 停止测试环境Docker容器
test-docker-stop:
	docker stop ca-test && docker rm ca-test
	@echo "测试环境已停止并删除"

# 显示帮助信息
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@echo "  build                  清理并构建所有平台版本"
	@echo "  build-windows          构建Windows版本 (AMD64)"
	@echo "  build-windows-arm64    构建Windows版本 (ARM64)"
	@echo "  build-darwin           构建macOS版本 (AMD64)"
	@echo "  build-darwin-arm64     构建macOS版本 (ARM64)"
	@echo "  build-linux            构建Linux版本 (AMD64)"
	@echo "  build-linux-arm64      构建Linux版本 (ARM64)"
	@echo "  build-linux-arm        构建Linux版本 (ARM)"
	@echo "  build-local            本地构建"
	@echo "  clean                  清理构建文件"
	@echo "  test                   运行测试"
	@echo "  test-cert              生成测试证书"
	@echo "  test-docker-build      构建测试环境Docker镜像"
	@echo "  test-docker-run        运行测试环境Docker容器"
	@echo "  test-docker-stop       停止测试环境Docker容器"
	@echo "  help                   显示此帮助信息"