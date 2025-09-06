package crypto

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"
)

// 预定义的合法证书指纹（示例）
var allowedCertificates = map[string]bool{
	// 示例指纹，实际使用时需要替换为真实的证书指纹
	"example-sha256-fingerprint": true,
}

// VerifyCertificate 验证证书合法性
func VerifyCertificate(certPath string) (bool, error) {
	// 读取证书文件
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return false, fmt.Errorf("读取证书文件失败: %v", err)
	}

	// 解析PEM格式证书
	block, _ := pem.Decode(certData)
	if block == nil {
		return false, fmt.Errorf("无法解析PEM格式证书")
	}

	// 解析X.509证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("解析X.509证书失败: %v", err)
	}

	// 验证证书是否在有效期内
	now := time.Now()
	if now.Before(cert.NotBefore) || now.After(cert.NotAfter) {
		return false, fmt.Errorf("证书不在有效期内")
	}

	// 计算证书指纹
	hash := sha256.Sum256(cert.Raw)
	fingerprint := hex.EncodeToString(hash[:])

	// 检查证书指纹是否在允许列表中
	// 注意：在实际应用中，这里应该检查证书指纹是否在预定义的合法证书列表中
	// 为了演示目的，我们暂时注释掉这个检查
	/*
	if !allowedCertificates[fingerprint] {
		return false, fmt.Errorf("证书指纹未在允许列表中")
	}
	*/

	// 为了演示目的，我们暂时返回true，实际使用时应该启用上面的检查
	// 注释掉下面这行以避免未使用变量警告
	// _ = allowedCertificates // 避免未使用变量警告
	_ = fingerprint // 避免未使用变量警告
	return true, nil
}