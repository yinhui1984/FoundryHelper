package utils

import (
	"bufio"
	"fmt"
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
func RunCommandLine2(shell, cmd string) (string, error) {
	outBytes, err := exec.Command(shell, "-c", cmd).CombinedOutput()
	return string(outBytes), err
}

// 运行命令并流式输出（其支持通配符管道符等特殊字符）， 当执行的命令比较耗时，并且命令一边执行一边输出时，可以使用此方法
func RunCommandLine3(shell, cmd string) error {
	command := exec.Command(shell, "-c", cmd)
	stdout, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return command.Wait()
}
