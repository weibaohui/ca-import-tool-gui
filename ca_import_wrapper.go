package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
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
