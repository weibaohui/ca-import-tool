package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/weibaohui/ca-import-tool/crypto"
	"github.com/weibaohui/ca-import-tool/docker"
	"github.com/weibaohui/ca-import-tool/platform"
)

var (
	dockerHost string
	force      bool
	version    = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "ca-import-tool [flags] <certificate-file>",
	Short: "CA证书自动导入工具",
	Long:  `一个跨平台的命令行工具，用于将CA证书导入系统信任库并配置Docker证书目录。`,
	Args:  cobra.ExactArgs(1),
	RunE:  run,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&dockerHost, "docker-host", "d", "", "指定Docker镜像仓库域名")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "强制覆盖已存在的证书")
	rootCmd.Flags().BoolP("version", "v", false, "显示版本信息")

	// 版本信息处理
	rootCmd.Flags().BoolP("help", "h", false, "显示帮助信息")
}

func run(cmd *cobra.Command, args []string) error {
	// 检查版本标志
	if ok, _ := cmd.Flags().GetBool("version"); ok {
		fmt.Printf("CA证书自动导入工具版本: %s\n", version)
		return nil
	}

	certPath := args[0]

	// 验证证书
	valid, err := crypto.VerifyCertificate(certPath)
	if err != nil {
		return fmt.Errorf("证书验证失败: %v", err)
	}
	if !valid {
		return fmt.Errorf("证书验证未通过")
	}

	// 检测操作系统
	osType := platform.DetectOS()
	fmt.Printf("检测到的操作系统: %s\n", osType)

	// 导入证书
	err = platform.ImportCertificate(certPath, osType, force)
	if err != nil {
		// 检查是否是因为不匹配的构建平台导致的错误
		return fmt.Errorf("证书导入失败: %v", err)
	}

	fmt.Println("证书已成功导入系统信任库")

	// 配置Docker证书
	if dockerHost != "" {
		err = docker.ConfigureDocker(certPath, dockerHost, force)
		if err != nil {
			return fmt.Errorf("Docker证书配置失败: %v", err)
		}
		fmt.Println("Docker证书配置完成")
	}

	fmt.Println("所有操作已完成，请重启浏览器和Docker服务使配置生效")

	return nil
}
