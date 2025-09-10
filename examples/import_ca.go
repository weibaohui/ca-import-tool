// Package main demonstrates how to use the caimport package
package main

import (
	"fmt"
	"log"

	"github.com/weibaohui/ca-import-tool/pkg/caimport"
)

func main() {
	// 示例：验证证书并导入
	// 注意：这需要一个真实的证书文件路径

	// 验证证书
	certPath := "/path/to/your/certificate.crt" // 替换为实际的证书路径
	valid, err := caimport.VerifyCertificate(certPath)
	if err != nil {
		log.Fatalf("证书验证失败: %v", err)
	}

	if !valid {
		log.Fatal("证书验证未通过")
	}

	// 导入证书
	options := caimport.ImportOptions{
		CertPath:   certPath,
		DockerHost: "registry.example.com", // 可选，如果需要为Docker配置证书
		Force:      false,                  // 是否强制覆盖已存在的证书
	}

	err = caimport.ImportCA(options)
	if err != nil {
		log.Fatalf("证书导入失败: %v", err)
	}

	fmt.Println("证书导入成功")
}
