//go:build linux

package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ImportCertificateLinux 将CA证书导入Linux系统信任库
func ImportCertificateLinux(certPath string, force bool) error {
	// 检查是Debian/Ubuntu还是CentOS/RHEL系统
	var certDir string
	var updateCmd *exec.Cmd
	
	// 检查是否存在Debian/Ubuntu的证书目录
	if _, err := os.Stat("/usr/local/share/ca-certificates/"); err == nil {
		// Debian/Ubuntu系统
		certDir = "/usr/local/share/ca-certificates/"
		updateCmd = exec.Command("sudo", "update-ca-certificates")
	} else if _, err := os.Stat("/etc/pki/ca-trust/source/anchors/"); err == nil {
		// CentOS/RHEL系统
		certDir = "/etc/pki/ca-trust/source/anchors/"
		updateCmd = exec.Command("sudo", "update-ca-trust")
	} else {
		return fmt.Errorf("不支持的Linux发行版或证书目录不存在")
	}
	
	// 复制证书文件
	certName := filepath.Base(certPath)
	targetPath := filepath.Join(certDir, certName)
	
	// 使用sudo复制文件
	cpCmd := exec.Command("sudo", "cp", certPath, targetPath)
	output, err := cpCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "sudo: a password is required") {
			return fmt.Errorf("需要管理员权限，请输入密码")
		}
		return fmt.Errorf("复制证书文件失败: %v, 输出: %s", err, string(output))
	}
	
	// 更新证书库
	output, err = updateCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("更新证书库失败: %v, 输出: %s", err, string(output))
	}
	
	return nil
}

// ImportCertificate 是Linux平台的证书导入函数
func ImportCertificate(certPath string, osType string, force bool) error {
	return ImportCertificateLinux(certPath, force)
}