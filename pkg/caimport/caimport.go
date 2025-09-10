// Package caimport 提供CA证书导入功能
package caimport

import (
	"fmt"
	
	"ca-import-tool/crypto"
	"ca-import-tool/docker"
	"ca-import-tool/platform"
)

// ImportOptions 导入选项
type ImportOptions struct {
	// CertPath 证书文件路径
	CertPath string
	
	// DockerHost Docker镜像仓库域名
	DockerHost string
	
	// Force 是否强制覆盖已存在的证书
	Force bool
}

// ImportCA 导入CA证书的统一入口
func ImportCA(options ImportOptions) error {
	// 验证证书
	valid, err := VerifyCertificate(options.CertPath)
	if err != nil {
		return fmt.Errorf("证书验证失败: %v", err)
	}
	if !valid {
		return fmt.Errorf("证书验证未通过")
	}

	// 检测操作系统
	osType := platform.DetectOS()
	// fmt.Printf("检测到的操作系统: %s\n", osType) // 移除打印，保持API清洁

	// 导入证书
	err = platform.ImportCertificate(options.CertPath, osType, options.Force)
	if err != nil {
		// 检查是否是因为不匹配的构建平台导致的错误
		return fmt.Errorf("证书导入失败: %v", err)
	}

	// fmt.Println("证书已成功导入系统信任库") // 移除打印，保持API清洁

	// 配置Docker证书
	if options.DockerHost != "" {
		err = docker.ConfigureDocker(options.CertPath, options.DockerHost, options.Force)
		if err != nil {
			return fmt.Errorf("Docker证书配置失败: %v", err)
		}
		// fmt.Println("Docker证书配置完成") // 移除打印，保持API清洁
	}

	// fmt.Println("所有操作已完成，请重启浏览器和Docker服务使配置生效") // 移除打印，保持API清洁

	return nil
}

// VerifyCertificate 验证证书合法性
func VerifyCertificate(certPath string) (bool, error) {
	return crypto.VerifyCertificate(certPath)
}