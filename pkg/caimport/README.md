# CA Import Package

这是一个用于导入CA证书的Go包，提供了将CA证书导入系统信任库并配置Docker证书目录的功能。

## 功能特性

- 跨平台支持（Windows、macOS、Linux）
- 自动识别操作系统类型
- 将CA证书导入系统信任库
- 为Docker配置证书目录
- 内置证书验证机制

## 安装

```bash
go get ca-import-tool/pkg/caimport
```

## 使用方法

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "ca-import-tool/pkg/caimport"
)

func main() {
    // 验证证书
    valid, err := caimport.VerifyCertificate("/path/to/certificate.crt")
    if err != nil {
        log.Fatalf("证书验证失败: %v", err)
    }
    
    if !valid {
        log.Fatal("证书验证未通过")
    }
    
    // 导入证书
    options := caimport.ImportOptions{
        CertPath:   "/path/to/certificate.crt",
        DockerHost: "",
        Force:      false,
    }
    
    err = caimport.ImportCA(options)
    if err != nil {
        log.Fatalf("证书导入失败: %v", err)
    }
    
    fmt.Println("证书导入成功")
}
```

### 完整流程使用

```go
package main

import (
    "fmt"
    "log"
    
    "ca-import-tool/pkg/caimport"
)

func main() {
    options := caimport.ImportOptions{
        CertPath:   "/path/to/certificate.crt",
        DockerHost: "registry.example.com",
        Force:      false,
    }
    
    // 验证证书
    valid, err := caimport.VerifyCertificate(options.CertPath)
    if err != nil || !valid {
        log.Fatalf("证书验证失败: %v", err)
    }
    
    // 导入证书和配置Docker
    err = caimport.ImportCA(options)
    if err != nil {
        log.Fatalf("操作失败: %v", err)
    }
    
    fmt.Println("所有操作已完成")
}
```

## API参考

### ImportOptions

```go
type ImportOptions struct {
    CertPath   string  // 证书文件路径
    DockerHost string  // Docker镜像仓库域名
    Force      bool    // 是否强制覆盖已存在的证书
}
```

### ImportCA

```go
func ImportCA(options ImportOptions) error
```

导入CA证书的统一入口，完成所有操作。

### VerifyCertificate

```go
func VerifyCertificate(certPath string) (bool, error)
```

验证证书合法性。

## 注意事项

1. 需要管理员权限才能执行系统级操作
2. 仅支持预定义的CA证书（通过证书指纹验证）
3. 在Linux系统上，支持Debian/Ubuntu和CentOS/RHEL发行版
4. 在Windows系统上，需要以管理员身份运行
5. 在macOS系统上，需要输入管理员密码