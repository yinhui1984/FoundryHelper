package utils

import (
	"os/exec"
	"strings"
)

// 其不支持通配符管道符等特殊字符
func RunCommandLine(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	outBytes, err := exec.Command(parts[0], parts[1:]...).CombinedOutput()
	return string(outBytes), err
}

// 其支持通配符管道符等特殊字符
func RunCommandLine2(cmd string) (string, error) {
	outBytes, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(outBytes), err
}
