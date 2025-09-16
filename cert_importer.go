package main

import (
	"runtime"
)

// CertificateImporter 证书导入接口
type CertificateImporter interface {
	// Import 执行证书导入操作
	Import(params ImportParams) ImportResult
	
	// Validate 验证证书文件
	Validate(filePath string) (bool, error)
	
	// List 列出已导入的证书
	List() []CertificateInfo
}

// NewCertificateImporter 创建证书导入器实例
func NewCertificateImporter() CertificateImporter {
	switch runtime.GOOS {
	case "darwin":
		return &MacCertificateImporter{}
	case "linux":
		return &LinuxCertificateImporter{}
	case "windows":
		return &WindowsCertificateImporter{}
	default:
		// 默认使用macOS实现
		return &MacCertificateImporter{}
	}
}