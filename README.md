# CA证书自动导入工具

一个跨平台的命令行工具，用于将CA证书导入系统信任库并配置Docker证书目录。

## 功能特性

- 跨平台支持：Windows、macOS、主流Linux发行版
- 自动识别操作系统类型
- 将CA证书导入系统信任库
- 为Docker配置证书目录
- 提供权限检查和错误处理机制
- 安全性保障：只允许导入预定义的CA证书

## 安装

### 从源码编译

```bash
# 克隆项目
git clone <repository-url>
cd ca-import-tool

# 编译
go build -o ca-import-tool

# 或者交叉编译其他平台
# Windows: GOOS=windows GOARCH=amd64 go build -o ca-import-tool.exe
# macOS: GOOS=darwin GOARCH=amd64 go build -o ca-import-tool-mac
# Linux: GOOS=linux GOARCH=amd64 go build -o ca-import-tool-linux
```

## 使用方法

```bash
# 基本用法：导入证书到系统信任库
./ca-import-tool ca.crt

# 同时为Docker配置证书
./ca-import-tool -d registry.example.com ca.crt

# 强制覆盖已存在的证书
./ca-import-tool -f ca.crt

# 查看帮助信息
./ca-import-tool -h

# 查看版本信息
./ca-import-tool -v
```

## 命令行参数

```
Usage:
  ca-import-tool [flags] <certificate-file>

Flags:
  -d, --docker-host string   指定Docker镜像仓库域名
  -f, --force                强制覆盖已存在的证书
  -h, --help                 显示帮助信息
  -v, --version              显示版本信息
```

## 平台支持

### Windows

使用`certutil`命令将证书导入Windows系统信任库：
```
certutil -addstore -f "Root" your-ca.crt
```

### macOS

使用`security`命令将证书导入macOS系统钥匙串：
```
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain your-ca.crt
```

### Linux

根据发行版不同使用相应命令：

**Debian/Ubuntu:**
```
sudo cp your-ca.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

**CentOS/RHEL:**
```
sudo cp your-ca.crt /etc/pki/ca-trust/source/anchors/
sudo update-ca-trust
```

## Docker证书配置

为Docker配置独立的证书信任目录：
```
sudo mkdir -p /etc/docker/certs.d/<registry-domain>/
sudo cp your-ca.crt /etc/docker/certs.d/<registry-domain>/ca.crt
```

## 安全说明

1. 所有系统级操作都需要管理员权限
2. 工具启动时会检查权限并在需要时提示用户
3. 证书验证模块确保只导入预定义的CA证书
4. 限制用户不能指定任意证书文件导入

## 许可证

[MIT License](LICENSE)