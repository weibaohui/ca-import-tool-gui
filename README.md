# CA证书导入工具

## 关于项目

这是一个基于 Wails 框架开发的桌面应用程序，专门用于在 macOS 系统上导入 CA 证书到系统钥匙串中。该工具提供了一个直观的图形用户界面，简化了证书导入流程，并处理了需要管理员权限的操作。

主要功能：
- 选择并验证 PEM/DER 格式的证书文件
- 以管理员权限将证书导入到系统钥匙串
- 显示导入操作的详细结果和日志

## 技术栈

- **后端**: Go 1.23 + Wails v2.10.2
- **前端**: React 18 + TypeScript + Ant Design 5
- **构建工具**: Vite 3
- **系统集成**: macOS `security` 命令行工具

## 系统要求

- macOS 操作系统（需要支持 osascript 和 security 命令）
- Go 1.23 或更高版本（用于开发和构建）
- Node.js（用于前端开发）

## 开发环境搭建

1. 安装 Go 1.23+：
   ```bash
   # 使用 Homebrew 安装 Go
   brew install go
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

## 安全说明

- 该工具使用 macOS 系统原生命令进行证书导入操作
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
└── wails.json          # Wails 项目配置
```

## 核心功能模块

### 证书导入 (ImportCertificate)
使用 macOS 的 `security add-trusted-cert` 命令将证书添加到系统钥匙串。

### 证书验证 (ValidateCertificate)
验证证书文件格式是否正确（支持 PEM 和 DER 格式）。

### 文件选择 (SelectCertificateFile)
打开系统文件选择对话框，让用户选择证书文件。

### 证书列表 (ListCertificates)
获取已导入证书的列表信息。

## 注意事项

1. 该工具仅支持 macOS 系统
2. 导入证书需要管理员权限，系统会提示输入密码
3. 请确保选择的证书文件格式正确
4. 如果导入过程中取消了权限授权，操作会被视为成功（这是 macOS 的安全机制）

## 许可证

本项目基于 MIT 许可证发布。