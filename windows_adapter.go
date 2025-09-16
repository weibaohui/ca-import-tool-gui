package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// WindowsCertificateImporter Windows证书导入适配器
type WindowsCertificateImporter struct{}

// Import 执行Windows证书导入操作
func (w *WindowsCertificateImporter) Import(params ImportParams) ImportResult {
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
	valid, err := w.Validate(params.FilePath)
	if !valid {
		result.Message = "证书验证失败"
		result.Log = fmt.Sprintf("错误: %v", err)
		return result
	}

	// 在Windows上使用certutil命令导入证书
	// certutil -addstore -f "ROOT" certificate.cer
	cmd := exec.Command("certutil", "-addstore", "-f", "ROOT", params.FilePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// 检查是否是因为权限问题导致的错误
		errorMsg := string(output)
		if strings.Contains(errorMsg, "拒绝访问") || strings.Contains(errorMsg, "Access is denied") {
			result.Success = false
			result.Message = "证书导入失败，需要管理员权限"
			result.Log = fmt.Sprintf("权限错误: %v, 输出: %s", err, errorMsg)
			return result
		}

		result.Success = false
		result.Message = "证书导入失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, errorMsg)
		return result
	}

	result.Success = true
	result.Message = "证书导入成功"
	result.Log = fmt.Sprintf("证书已成功导入到系统证书存储中，输出: %s", string(output))
	return result
}

// Validate 验证证书文件
func (w *WindowsCertificateImporter) Validate(filePath string) (bool, error) {
	return ValidateCertificate(filePath)
}

// List 列出已导入的证书
func (w *WindowsCertificateImporter) List() []CertificateInfo {
	// 实现Windows平台的证书列表功能
	// 这里返回模拟数据，实际实现可能需要通过PowerShell或certutil命令查询
	return ListCertificates()
}