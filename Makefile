# 应用程序名称和版本信息
APP_NAME := caimport
VERSION := 1.0.0
GIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S)
IDENTIFIER := www.mine.app
PKG_OUTPUT := build/$(APP_NAME)-$(VERSION)

# 编译标记
LDFLAGS := -ldflags="-X 'main.AppName=$(APP_NAME)' -X 'main.Version=$(VERSION)' -X 'main.GitHash=$(GIT_HASH)' -X 'main.BuildTime=$(BUILD_TIME)'"

# Wails 命令
WAILS := wails

# 默认目标
.PHONY: all
all: clean build package

# 清理旧的构建文件
.PHONY: clean
clean:
	@echo "Cleaning build directory..."
	rm -rf build

# 生成静态文件（如需要）
.PHONY: static-generated
static-generated:
	statik -src=./configs

# 运行开发模式
.PHONY: dev
dev:
	$(WAILS) dev -s

# 各个平台和架构的构建目标
.PHONY: build
build: mac_amd64 mac_arm64  windows_amd64 windows_arm64 linux_amd64 linux_arm64

# 构建 Linux 版本（仅Linux平台）
.PHONY: build-linux
build-linux: linux_amd64 linux_arm64

# 构建 macOS 版本
.PHONY: mac_amd64
mac_amd64:
	@echo "Building macOS AMD64 version with Git Hash: $(GIT_HASH)"
	GOOS=darwin GOARCH=amd64 $(WAILS) build -platform darwin/amd64 $(LDFLAGS) -o $(APP_NAME)-macOS-AMD64
	@echo "Moving output to build directory..."
	@mkdir -p build
	@rm -rf "build/$(APP_NAME)-macOS-AMD64.app"
	@if [ -d "build/bin/$(APP_NAME).app" ]; then \
		mv "build/bin/$(APP_NAME).app" "build/$(APP_NAME)-macOS-AMD64.app"; \
	elif [ -f "build/bin/$(APP_NAME)" ]; then \
		mv "build/bin/$(APP_NAME)" "build/$(APP_NAME)-macOS-AMD64"; \
	fi
	# macOS AMD64 应用签名
	@if [ -d "build/$(APP_NAME)-macOS-AMD64.app" ]; then \
		echo "Signing macOS AMD64 app..."; \
		codesign --force --deep -s - "build/$(APP_NAME)-macOS-AMD64.app"; \
	fi

.PHONY: mac_arm64
mac_arm64:
	@echo "Building macOS ARM64 version with Git Hash: $(GIT_HASH)"
	GOOS=darwin GOARCH=arm64 $(WAILS) build -platform darwin/arm64 $(LDFLAGS) -o $(APP_NAME)-macOS-ARM64
	@echo "Moving output to build directory..."
	@mkdir -p build
	@rm -rf "build/$(APP_NAME)-macOS-ARM64.app"
	@if [ -d "build/bin/$(APP_NAME).app" ]; then \
		mv "build/bin/$(APP_NAME).app" "build/$(APP_NAME)-macOS-ARM64.app"; \
	elif [ -f "build/bin/$(APP_NAME)" ]; then \
		mv "build/bin/$(APP_NAME)" "build/$(APP_NAME)-macOS-ARM64"; \
	fi
	# macOS ARM64 应用签名
	@if [ -d "build/$(APP_NAME)-macOS-ARM64.app" ]; then \
		echo "Signing macOS ARM64 app..."; \
		codesign --force --deep -s - "build/$(APP_NAME)-macOS-ARM64.app"; \
	fi

# 构建 Windows 版本
.PHONY: windows_amd64
windows_amd64:
	@echo "Building Windows AMD64 version with Git Hash: $(GIT_HASH)"
	GOOS=windows GOARCH=amd64 $(WAILS) build -platform windows/amd64 $(LDFLAGS) -o $(APP_NAME)-Windows-AMD64.exe
	@echo "Moving output to build directory..."
	@mkdir -p build
	@if [ -f "build/bin/$(APP_NAME).exe" ]; then \
		mv "build/bin/$(APP_NAME).exe" "build/$(APP_NAME)-Windows-AMD64.exe"; \
	fi

.PHONY: windows_arm64
windows_arm64:
	@echo "Building Windows ARM64 version with Git Hash: $(GIT_HASH)"
	GOOS=windows GOARCH=arm64 $(WAILS) build -platform windows/arm64 $(LDFLAGS) -o $(APP_NAME)-Windows-ARM64.exe
	@echo "Moving output to build directory..."
	@mkdir -p build
	@if [ -f "build/bin/$(APP_NAME).exe" ]; then \
		mv "build/bin/$(APP_NAME).exe" "build/$(APP_NAME)-Windows-ARM64.exe"; \
	fi

# 构建 Linux 版本
.PHONY: linux_amd64
linux_amd64:
	@echo "Building Linux AMD64 version with Git Hash: $(GIT_HASH)"
	GOOS=linux GOARCH=amd64 $(WAILS) build -platform linux/amd64 $(LDFLAGS) -o $(APP_NAME)-Linux-AMD64
	@echo "Moving output to build directory..."
	@mkdir -p build
	@if [ -f "build/bin/$(APP_NAME)" ]; then \
		mv "build/bin/$(APP_NAME)" "build/$(APP_NAME)-Linux-AMD64"; \
	fi

.PHONY: linux_arm64
linux_arm64:
	@echo "Building Linux ARM64 version with Git Hash: $(GIT_HASH)"
	GOOS=linux GOARCH=arm64 $(WAILS) build -platform linux/arm64 $(LDFLAGS) -o $(APP_NAME)-Linux-ARM64
	@echo "Moving output to build directory..."
	@mkdir -p build
	@if [ -f "build/bin/$(APP_NAME)" ]; then \
		mv "build/bin/$(APP_NAME)" "build/$(APP_NAME)-Linux-ARM64"; \
	fi

# 打包目标：打包所有平台和架构的应用
.PHONY: package
package: package_mac package_windows package_linux

# 打包 Linux 版本（仅Linux平台）
.PHONY: package-linux
package-linux: build-linux package_linux

# macOS 打包
.PHONY: package_mac
package_mac: mac_amd64 mac_arm64 
	@echo "Packaging macOS versions..."
	hdiutil create -volname "$(APP_NAME) AMD64" -srcfolder build/$(APP_NAME)-macOS-AMD64 -ov -format UDZO $(PKG_OUTPUT)-macOS-AMD64.dmg
	hdiutil create -volname "$(APP_NAME) ARM64" -srcfolder build/$(APP_NAME)-macOS-ARM64 -ov -format UDZO $(PKG_OUTPUT)-macOS-ARM64.dmg

# Windows 打包
.PHONY: package_windows
package_windows: windows_amd64 windows_arm64
	@echo "Packaging Windows versions..."
	zip -j $(PKG_OUTPUT)-Windows-AMD64.zip build/$(APP_NAME)-Windows-AMD64.exe
	zip -j $(PKG_OUTPUT)-Windows-ARM64.zip build/$(APP_NAME)-Windows-ARM64.exe

# Linux 打包
.PHONY: package_linux
package_linux: linux_amd64 linux_arm64
	@echo "Packaging Linux versions..."
	tar -czvf $(PKG_OUTPUT)-Linux-AMD64.tar.gz -C build $(APP_NAME)-Linux-AMD64
	tar -czvf $(PKG_OUTPUT)-Linux-ARM64.tar.gz -C build $(APP_NAME)-Linux-ARM64

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all              : Clean, build, and package the app for all platforms and architectures"
	@echo "  build            : Build the app for all platforms and architectures with version and Git hash"
	@echo "  build-linux      : Build the app for Linux platforms (AMD64 and ARM64)"
	@echo "  clean            : Clean the build directory"
	@echo "  mac_amd64        : Build macOS AMD64 version"
	@echo "  mac_arm64        : Build macOS ARM64 version"
	@echo "  windows_amd64    : Build Windows AMD64 version"
	@echo "  windows_arm64    : Build Windows ARM64 version"
	@echo "  linux_amd64      : Build Linux AMD64 version"
	@echo "  linux_arm64      : Build Linux ARM64 version"
	@echo "  package          : Package the built apps into distributable files"
	@echo "  package-linux    : Build and package Linux apps"
	@echo "  package_mac      : Package macOS apps"
	@echo "  package_windows  : Package Windows apps"
	@echo "  package_linux    : Package Linux apps"