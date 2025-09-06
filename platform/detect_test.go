package platform

import (
	"runtime"
	"testing"
)

func TestDetectOS(t *testing.T) {
	osType := DetectOS()
	
	// 检查返回的操作系统类型是否有效
	validOS := []string{"windows", "darwin", "linux"}
	found := false
	for _, os := range validOS {
		if osType == os {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("DetectOS() returned invalid OS type: %s", osType)
	}
	
	// 检查是否与runtime.GOOS一致
	if osType != runtime.GOOS {
		t.Errorf("DetectOS() = %s, want %s", osType, runtime.GOOS)
	}
}