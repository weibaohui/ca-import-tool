package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConfigureDocker 为Docker配置证书目录
func ConfigureDocker(certPath string, registry string, force bool) error {
	// Docker证书目录路径
	dockerCertsDir := fmt.Sprintf("/etc/docker/certs.d/%s", registry)
	
	// 创建目录
	mkdirCmd := exec.Command("sudo", "mkdir", "-p", dockerCertsDir)
	output, err := mkdirCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "sudo: a password is required") {
			return fmt.Errorf("需要管理员权限，请输入密码")
		}
		return fmt.Errorf("创建Docker证书目录失败: %v, 输出: %s", err, string(output))
	}
	
	// 复制证书文件
	targetPath := filepath.Join(dockerCertsDir, "ca.crt")
	
	// 检查文件是否已存在且不强制覆盖
	if !force {
		if _, err := os.Stat(targetPath); err == nil {
			return fmt.Errorf("证书文件已存在，使用 -f 参数强制覆盖")
		}
	}
	
	cpCmd := exec.Command("sudo", "cp", certPath, targetPath)
	output, err = cpCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "sudo: a password is required") {
			return fmt.Errorf("需要管理员权限，请输入密码")
		}
		return fmt.Errorf("复制证书到Docker目录失败: %v, 输出: %s", err, string(output))
	}
	
	return nil
}