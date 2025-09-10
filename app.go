package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ImportCertificate 执行证书导入操作
func (a *App) ImportCertificate(params ImportParams) (ImportResult, error) {
	result := ImportCertificate(params)
	return result, nil
}

// SelectCertificateFile 打开文件选择对话框并返回选中的文件路径
func (a *App) SelectCertificateFile() (string, error) {
	selectedFile, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择证书文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "证书文件 (*.pem, *.cer, *.crt, *.der)",
				Pattern:     "*.pem;*.cer;*.crt;*.der",
			},
		},
	})

	if err != nil {
		return "", err
	}

	return selectedFile, nil
}

// ValidateCertificate 验证证书文件
func (a *App) ValidateCertificate(filePath string) (bool, error) {
	valid, err := ValidateCertificate(filePath)
	return valid, err
}

// ListCertificates 列出已导入的证书
func (a *App) ListCertificates() ([]CertificateInfo, error) {
	certs := ListCertificates()
	return certs, nil
}
