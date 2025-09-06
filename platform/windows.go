//go:build windows

package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

// ImportCertificateWindows 将CA证书导入Windows系统信任库
func ImportCertificateWindows(certPath string, force bool) error {
	// Windows使用certutil命令导入证书
	cmd := exec.Command("certutil", "-addstore", "-f", "Root", certPath)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		if strings.Contains(string(output), "拒绝访问") {
			return fmt.Errorf("权限不足，请以管理员身份运行此工具")
		}
		return fmt.Errorf("导入证书失败: %v, 输出: %s", err, string(output))
	}
	
	return nil
}

// ImportCertificate 是Windows平台的证书导入函数
func ImportCertificate(certPath string, osType string, force bool) error {
	return ImportCertificateWindows(certPath, force)
}