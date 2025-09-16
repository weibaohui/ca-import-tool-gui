# CA证书导入工具

## 关于项目

这是一个基于 Wails 框架开发的跨平台桌面应用程序，专门用于在 Windows、macOS 和 Linux 系统上导入 CA 证书。该工具提供了一个直观的图形用户界面，简化了证书导入流程，并处理了需要管理员权限的操作。

主要功能：
- 选择并验证 PEM/DER 格式的证书文件
- 以管理员权限将证书导入到系统证书存储中
- 显示导入操作的详细结果和日志
- 支持跨平台使用（Windows、macOS、Linux）

## 技术栈

- **后端**: Go 1.23 + Wails v2.10.2
- **前端**: React 18 + TypeScript + Ant Design 5
- **构建工具**: Vite 3
- **系统集成**: 
  - Windows: `certutil` 命令行工具
  - macOS: `security` 命令行工具
  - Linux: `update-ca-certificates` 或 `update-ca-trust` 命令行工具

## 系统要求

### Windows
- Windows 10 或更高版本
- 需要 PowerShell 支持

### macOS
- macOS 10.15 (Catalina) 或更高版本
- 需要支持 osascript 和 security 命令

### Linux
支持的主要发行版：
- Ubuntu 18.04 LTS 及以上版本
- Debian 10 及以上版本
- CentOS 7 及以上版本
- Red Hat Enterprise Linux 7 及以上版本
- Fedora 32 及以上版本
- openSUSE Leap 15 及以上版本
- Arch Linux (最新版本)
- Alpine Linux 3.12 及以上版本

## 开发环境搭建

1. 安装 Go 1.23+：
   ```bash
   # 使用 Homebrew 安装 Go (macOS)
   brew install go
   
   # 使用 Chocolatey 安装 Go (Windows)
   choco install golang
   
   # 使用 apt 安装 Go (Ubuntu/Debian)
   sudo apt install golang
   ```

2. 安装 Wails CLI：
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

3. 安装前端依赖：
   ```bash
   cd frontend
   npm install
   ```

## 本地开发

在项目根目录下运行以下命令启动开发模式：

```bash
wails dev
```

这将启动一个带有热重载功能的开发服务器。您可以在浏览器中访问 http://localhost:34115 进行开发调试。

## 构建生产版本

要构建可发布的生产版本，运行：

```bash
wails build
```

构建完成后，可在 `build/bin` 目录下找到可执行的二进制文件。

## 使用说明

1. 启动应用程序
2. 点击"选择证书文件"按钮选择要导入的证书文件（支持 .pem, .cer, .crt, .der 格式）
3. 点击"导入证书"按钮开始导入过程
4. 系统会弹出管理员权限授权对话框，输入密码确认操作
5. 查看导入结果和详细日志

## 跨平台支持详情

### Windows
- 使用 `certutil` 命令导入证书
- 自动检测管理员权限
- 需要用户确认UAC提示

### macOS
- 使用 `security add-trusted-cert` 命令将证书添加到系统钥匙串
- 通过 osascript 请求管理员权限

### Linux
- 自动检测系统类型并使用相应的证书管理工具：
  - Debian/Ubuntu/Alpine/openSUSE 系列：使用 `update-ca-certificates`
  - RHEL/CentOS/Fedora/Arch 系列：使用 `update-ca-trust`
- 需要 sudo 权限执行证书导入操作

## 安全说明

- 该工具使用各操作系统原生命令进行证书导入操作
- 所有操作都需要用户明确授权管理员权限
- 证书文件仅在本地处理，不会上传到任何外部服务器

## 项目结构

```
.
├── frontend/           # 前端 React 应用
│   ├── src/            # 前端源码
│   │   ├── components/ # React 组件
│   │   └── assets/     # 静态资源
│   └── wailsjs/        # Wails 自动生成的前端绑定代码
├── app.go              # Go 后端主逻辑
├── ca_import_wrapper.go # 证书导入核心功能实现
├── models.go           # 数据模型定义
├── main.go             # 应用程序入口
├── windows_adapter.go  # Windows平台适配器
├── mac_adapter.go      # macOS平台适配器
├── linux_adapter.go    # Linux平台适配器
├── windows_admin.go    # Windows管理员权限检测
├── non_windows_admin.go # 非Windows平台管理员权限检测
└── wails.json          # Wails 项目配置
```

## 核心功能模块

### 证书导入 (ImportCertificate)
根据操作系统平台使用相应的命令将证书添加到系统证书存储中。

### 证书验证 (ValidateCertificate)
验证证书文件格式是否正确（支持 PEM 和 DER 格式）。

### 文件选择 (SelectCertificateFile)
打开系统文件选择对话框，让用户选择证书文件。

### 证书列表 (ListCertificates)
获取已导入证书的列表信息。

### 系统信息 (GetSystemInfo)
获取当前系统信息，包括操作系统类型、架构、Go版本等。

## 注意事项

1. 导入证书需要管理员权限，系统会提示输入密码或确认UAC对话框
2. 请确保选择的证书文件格式正确
3. 在Linux系统上，可能需要根据发行版调整证书管理命令
4. 如果导入过程中取消了权限授权，操作会被视为失败

## 许可证

本项目基于 MIT 许可证发布。