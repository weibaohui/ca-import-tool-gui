package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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
	// 处理Windows系统的编码问题
	decodedOutput, decodeErr := decodeWindowsOutput(output)
	if decodeErr != nil {
		decodedOutput = string(output) // 如果解码失败，使用原始输出
	}

	if err != nil {
		result.Success = false
		result.Message = "证书导入失败"
		result.Log = fmt.Sprintf("执行错误: %v, 输出: %s", err, decodedOutput)
		return result
	}

	result.Success = true
	result.Message = "证书导入成功"
	result.Log = fmt.Sprintf("证书已成功导入到系统证书存储中，输出: %s", decodedOutput)
	return result
}

// decodeWindowsOutput 处理Windows系统的编码问题
func decodeWindowsOutput(output []byte) (string, error) {
	// 尝试使用GBK解码
	reader := transform.NewReader(bytes.NewReader(output), simplifiedchinese.GBK.NewDecoder())
	decoded, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
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
