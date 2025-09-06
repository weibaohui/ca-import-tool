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

## 测试环境

项目提供了一个完整的测试环境，用于验证CA证书导入工具的功能：

1. 使用自签名证书的测试服务器
2. 对比证书导入前后的浏览器信任状态

详细说明请参考 [测试环境说明](test/README.md)。

## 许可证

[MIT License](LICENSE)
