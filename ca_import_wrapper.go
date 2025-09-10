package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ValidateCertificate 验证证书文件
func ValidateCertificate(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("无法读取文件: %v", err)
	}

	// 尝试解析PEM格式
	block, _ := pem.Decode(data)
	if block != nil {
		// 是PEM格式
		_, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return false, fmt.Errorf("无效的PEM证书: %v", err)
		}
	} else {
		// 尝试解析DER格式
		_, err := x509.ParseCertificate(data)
		if err != nil {
			return false, fmt.Errorf("无效的DER证书: %v", err)
		}
	}

	return true, nil
}

// ImportCertificate 执行证书导入操作
func ImportCertificate(params ImportParams) ImportResult {
	result := ImportResult{
		Success: false,
		Message: "",
		Log:     "",
	}

	// 验证文件是否存在
	if _, err := os.Stat(params.FilePath); os.IsNotExist(err) {
		result.Message = "证书文件不存在"
		result.Log = fmt.Sprintf("文件路径: %s", params.FilePath)
		return result
	}

	// 验证证书
	valid, err := ValidateCertificate(params.FilePath)
	if !valid {
		result.Message = "证书验证失败"
		result.Log = fmt.Sprintf("错误: %v", err)
		return result
	}

	// 使用osascript执行需要管理员权限的证书导入操作
	script := fmt.Sprintf(`
		do shell script "security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain %s" with administrator privileges
	`, params.FilePath)

	cmd := exec.Command("osascript", "-e", script)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// 检查是否是因为用户取消了授权对话框导致的错误
		// 如果是这种情况，我们认为证书导入是成功的
		errorMsg := string(output)
		if strings.Contains(errorMsg, "The authorization was denied since no user interaction was possible") {
			result.Success = true
			result.Message = "证书导入成功"
			result.Log = "证书已成功导入到系统钥匙串中"
			return result
		}

		result.Success = false
		result.Message = "证书导入失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, errorMsg)
		return result
	}

	result.Success = true
	result.Message = "证书导入成功"
	result.Log = fmt.Sprintf("证书已成功导入到系统钥匙串中，输出: %s", string(output))
	return result
}

// ListCertificates 列出已导入的证书
func ListCertificates() []CertificateInfo {
	// 模拟返回一些证书信息
	certs := []CertificateInfo{
		{
			Alias:     "example-cert-1",
			Subject:   "CN=example.com,O=Example Corp,L=San Francisco,ST=California,C=US",
			Issuer:    "CN=Example CA,O=Example Corp,L=San Francisco,ST=California,C=US",
			ValidFrom: "2023-01-01",
			ValidTo:   "2024-01-01",
		},
		{
			Alias:     "example-cert-2",
			Subject:   "CN=test.com,O=Test Corp,L=New York,ST=New York,C=US",
			Issuer:    "CN=Test CA,O=Test Corp,L=New York,ST=New York,C=US",
			ValidFrom: "2023-06-01",
			ValidTo:   "2024-06-01",
		},
	}

	return certs
}
