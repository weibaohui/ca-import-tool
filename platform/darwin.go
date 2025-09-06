//go:build darwin

package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

// ImportCertificateDarwin 将CA证书导入macOS系统钥匙串
func ImportCertificateDarwin(certPath string, force bool) error {
	// macOS使用security命令导入证书
	cmd := exec.Command("sudo", "security", "add-trusted-cert", "-d", "-r", "trustRoot", "-k", "/Library/Keychains/System.keychain", certPath)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		if strings.Contains(string(output), "sudo: a password is required") {
			return fmt.Errorf("需要管理员权限，请输入密码")
		}
		return fmt.Errorf("导入证书失败: %v, 输出: %s", err, string(output))
	}
	
	return nil
}

// ImportCertificate 是macOS平台的证书导入函数
func ImportCertificate(certPath string, osType string, force bool) error {
	return ImportCertificateDarwin(certPath, force)
}