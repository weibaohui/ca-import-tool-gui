package main

import (
	"context"
	"fmt"
	"runtime"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// SystemInfo 系统信息结构体
type SystemInfo struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	GoVersion string `json:"go_version"`
	AppName   string `json:"app_name"`
	AppVer    string `json:"app_ver"`
	IsAdmin   bool   `json:"is_admin"` // 添加管理员权限检测字段
}

// App struct
type App struct {
	ctx          context.Context
	certImporter CertificateImporter
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		certImporter: NewCertificateImporter(),
	}
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
	result := a.certImporter.Import(params)
	return result, nil
}

// SelectCertificateFiles 打开文件选择对话框并返回选中的文件路径列表
func (a *App) SelectCertificateFiles() ([]string, error) {
	selectedFiles, err := wailsruntime.OpenMultipleFilesDialog(a.ctx, wailsruntime.OpenDialogOptions{
		Title: "选择证书文件",
		Filters: []wailsruntime.FileFilter{
			{
				DisplayName: "证书文件 (*.pem, *.cer, *.crt, *.der)",
				Pattern:     "*.pem;*.cer;*.crt;*.der",
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return selectedFiles, nil
}

// ValidateCertificate 验证证书文件
func (a *App) ValidateCertificate(filePath string) (bool, error) {
	valid, err := a.certImporter.Validate(filePath)
	return valid, err
}

// ListCertificates 列出已导入的证书
func (a *App) ListCertificates() ([]CertificateInfo, error) {
	certs := a.certImporter.List()
	return certs, nil
}

// GetSystemInfo 获取系统信息
func (a *App) GetSystemInfo() SystemInfo {
	info := SystemInfo{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		AppName:   "CA证书导入工具",
		AppVer:    "1.0.0",
		IsAdmin:   false, // 默认值
	}

	// 仅在Windows平台检测管理员权限
	if runtime.GOOS == "windows" {
		admin, err := isWindowsAdmin()
		if err == nil {
			info.IsAdmin = admin
		}
	} else {
		// 非Windows平台默认认为有足够权限
		info.IsAdmin = true
	}

	return info
}
