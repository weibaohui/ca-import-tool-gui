package main

import (
	"fmt"
	"os"
	"os/exec"
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
	if _, err := os.Stat(params.FilePath); os.IsNotExist(err) {
		result.Message = "证书文件不存在"
		result.Log = fmt.Sprintf("文件路径: %s", params.FilePath)
		return result
	}

	// 验证证书
	valid, err := l.Validate(params.FilePath)
	if !valid {
		result.Message = "证书验证失败"
		result.Log = fmt.Sprintf("错误: %v", err)
		return result
	}

	// 在Linux上使用certutil或cp命令导入证书
	// 这里使用cp命令将证书复制到系统证书目录
	cmd := exec.Command("sudo", "cp", params.FilePath, "/usr/local/share/ca-certificates/")

	output, err := cmd.CombinedOutput()
	if err != nil {
		result.Success = false
		result.Message = "证书导入失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, string(output))
		return result
	}

	// 更新证书库
	updateCmd := exec.Command("sudo", "update-ca-certificates")
	updateOutput, updateErr := updateCmd.CombinedOutput()
	if updateErr != nil {
		result.Success = false
		result.Message = "证书库更新失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", updateErr, string(updateOutput))
		return result
	}

	result.Success = true
	result.Message = "证书导入成功"
	result.Log = fmt.Sprintf("证书已成功导入并更新证书库，输出: %s", string(updateOutput))
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