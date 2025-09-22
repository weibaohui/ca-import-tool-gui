package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// LinuxCertificateImporter Linux证书导入适配器
type LinuxCertificateImporter struct{}

// Import 执行Linux证书导入操作
func (l *LinuxCertificateImporter) Import(params ImportParams) ImportResult {
	result := ImportResult{
		Success: false,
		Message: "",
		Log:     "",
	}

	// 验证文件是否存在
	for _, filePath := range params.FilePaths {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			result.Message = "证书文件不存在"
			result.Log = fmt.Sprintf("文件路径: %s", filePath)
			return result
		}
	}

	// 验证证书
	for _, filePath := range params.FilePaths {
		valid, err := l.Validate(filePath)
		if !valid {
			result.Message = "证书验证失败"
			result.Log = fmt.Sprintf("错误: %v", err)
			return result
		}
	}

	// 检测Linux发行版并执行相应的证书导入操作
	return l.importCertificatesByDistribution(params)
}

// importCertificatesByDistribution 根据Linux发行版执行相应的证书导入操作
func (l *LinuxCertificateImporter) importCertificatesByDistribution(params ImportParams) ImportResult {
	result := ImportResult{
		Success: false,
		Message: "",
		Log:     "",
	}

	// 检测系统类型并执行相应的导入操作
	if l.isUpdateCaCertificatesSystem() {
		return l.importWithUpdateCaCertificates(params)
	} else if l.isUpdateCaTrustSystem() {
		return l.importWithUpdateCaTrust(params)
	} else {
		result.Message = "不支持的Linux发行版"
		result.Log = "无法识别系统的证书管理工具"
		return result
	}
}

// isUpdateCaCertificatesSystem 检测是否为使用update-ca-certificates的系统
// 包括Debian/Ubuntu/Alpine/openSUSE等
func (l *LinuxCertificateImporter) isUpdateCaCertificatesSystem() bool {
	// 检查update-ca-certificates命令是否存在
	_, err := exec.LookPath("update-ca-certificates")
	return err == nil
}

// isUpdateCaTrustSystem 检测是否为使用update-ca-trust的系统
// 包括RHEL/CentOS/Fedora/Arch等
func (l *LinuxCertificateImporter) isUpdateCaTrustSystem() bool {
	// 检查update-ca-trust命令是否存在
	_, err := exec.LookPath("update-ca-trust")
	return err == nil
}

// importWithUpdateCaCertificates 使用update-ca-certificates方式导入证书
func (l *LinuxCertificateImporter) importWithUpdateCaCertificates(params ImportParams) ImportResult {
	result := ImportResult{
		Success: false,
		Message: "",
		Log:     "",
	}

	var successCount int
	var logMessages []string

	// 复制所有证书文件到系统目录
	for _, filePath := range params.FilePaths {
		// 获取证书文件名
		fileName := filePath[strings.LastIndex(filePath, "/")+1:]

		// 复制证书到系统目录
		cmd := exec.Command("sudo", "cp", filePath, "/usr/local/share/ca-certificates/"+fileName)
		output, err := cmd.CombinedOutput()
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("证书 %s 复制失败", filePath)
			result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, string(output))
			return result
		}

		successCount++
		logMessages = append(logMessages, fmt.Sprintf("证书 %s 已复制到系统目录", filePath))
	}

	// 更新证书库（只需要执行一次）
	updateCmd := exec.Command("sudo", "update-ca-certificates")
	updateOutput, updateErr := updateCmd.CombinedOutput()
	if updateErr != nil {
		result.Success = false
		result.Message = "证书库更新失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", updateErr, string(updateOutput))
		return result
	}

	result.Success = true
	result.Message = fmt.Sprintf("成功导入 %d 个证书", successCount)
	result.Log = strings.Join(logMessages, "\n") + fmt.Sprintf("\n证书库已成功更新，输出: %s", string(updateOutput))
	return result
}

// importWithUpdateCaTrust 使用update-ca-trust方式导入证书
func (l *LinuxCertificateImporter) importWithUpdateCaTrust(params ImportParams) ImportResult {
	result := ImportResult{
		Success: false,
		Message: "",
		Log:     "",
	}

	var successCount int
	var logMessages []string

	// 复制所有证书文件到系统目录
	for _, filePath := range params.FilePaths {
		// 获取证书文件名
		fileName := filePath[strings.LastIndex(filePath, "/")+1:]

		// 复制证书到系统目录 (不同发行版路径可能不同，这里使用最常见的)
		// 对于RHEL/CentOS/Fedora/Rocky/AlmaLinux
		cmd := exec.Command("sudo", "cp", filePath, "/etc/pki/ca-trust/source/anchors/"+fileName)
		output, err := cmd.CombinedOutput()

		// 如果上述路径不存在，尝试Arch Linux的路径
		if err != nil {
			cmd = exec.Command("sudo", "cp", filePath, "/etc/ca-certificates/trust-source/anchors/"+fileName)
			output, err = cmd.CombinedOutput()
		}

		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("证书 %s 复制失败", filePath)
			result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, string(output))
			return result
		}

		successCount++
		logMessages = append(logMessages, fmt.Sprintf("证书 %s 已复制到系统目录", filePath))
	}

	// 更新证书库（只需要执行一次）
	updateCmd := exec.Command("sudo", "update-ca-trust")
	updateOutput, updateErr := updateCmd.CombinedOutput()
	if updateErr != nil {
		result.Success = false
		result.Message = "证书库更新失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", updateErr, string(updateOutput))
		return result
	}

	result.Success = true
	result.Message = fmt.Sprintf("成功导入 %d 个证书", successCount)
	result.Log = strings.Join(logMessages, "\n") + fmt.Sprintf("\n证书库已成功更新，输出: %s", string(updateOutput))
	return result
}

// Validate 验证证书文件
func (l *LinuxCertificateImporter) Validate(filePath string) (bool, error) {
	return ValidateCertificate(filePath)
}

// List 列出已导入的证书
func (l *LinuxCertificateImporter) List() []CertificateInfo {
	// 实现Linux平台的证书列表功能
	// 这里返回模拟数据，实际实现可能需要读取系统证书目录
	return ListCertificates()
}
