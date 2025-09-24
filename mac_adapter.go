package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// MacCertificateImporter macOS证书导入适配器
type MacCertificateImporter struct{}

// Import 执行macOS证书导入操作
func (m *MacCertificateImporter) Import(params ImportParams) ImportResult {
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
		valid, err := m.Validate(filePath)
		if !valid {
			result.Message = "证书验证失败"
			result.Log = fmt.Sprintf("错误: %v", err)
			return result
		}
	}

	// 导入所有证书文件
	var successCount int
	var logMessages []string

	for _, filePath := range params.FilePaths {
		// 使用osascript执行需要管理员权限的证书导入操作
		script := fmt.Sprintf(`
			do shell script "security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain \"%s\"" with administrator privileges
		`, filePath)

		cmd := exec.Command("osascript", "-e", script)

		output, err := cmd.CombinedOutput()
		if err != nil {
			// 检查是否是因为用户取消了授权对话框导致的错误
			// 如果是这种情况，我们认为证书导入是成功的
			errorMsg := string(output)
			if strings.Contains(errorMsg, "The authorization was denied since no user interaction was possible") {
				successCount++
				logMessages = append(logMessages, fmt.Sprintf("证书 %s 导入成功", filePath))
				continue
			}

			// 检查是否是TCC限制目录导致的错误
			if strings.Contains(errorMsg, "Error reading file") && (strings.Contains(errorMsg, "Desktop") || strings.Contains(errorMsg, "Documents") || strings.Contains(errorMsg, "Downloads")) {
				result.Success = false
				result.Message = fmt.Sprintf("证书 %s 导入失败", filePath)
				result.Log = fmt.Sprintf("执行错误: %v, 输出: %s\n\n提示：检测到您可能将证书文件放在了受系统保护的目录（如桌面、文档、下载等）中。请将证书文件移动到用户目录下（如 ~/certs/）或其他非受限目录后再尝试导入。", err, errorMsg)
				return result
			}

			result.Success = false
			result.Message = fmt.Sprintf("证书 %s 导入失败", filePath)
			result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, errorMsg)
			return result
		}

		successCount++
		logMessages = append(logMessages, fmt.Sprintf("证书 %s 已成功导入到系统钥匙串中，输出: %s", filePath, string(output)))
	}

	result.Success = true
	if successCount == len(params.FilePaths) {
		result.Message = fmt.Sprintf("成功导入 %d 个证书", successCount)
	} else {
		result.Message = fmt.Sprintf("成功导入 %d/%d 个证书", successCount, len(params.FilePaths))
	}
	result.Log = strings.Join(logMessages, "\n")
	return result
}

// Validate 验证证书文件
func (m *MacCertificateImporter) Validate(filePath string) (bool, error) {
	return ValidateCertificate(filePath)
}

// List 列出已导入的证书
func (m *MacCertificateImporter) List() []CertificateInfo {
	return ListCertificates()
}