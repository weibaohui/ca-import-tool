package platform

import (
	"runtime"
)

// DetectOS 检测操作系统类型
func DetectOS() string {
	return runtime.GOOS
}