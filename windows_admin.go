//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

// isWindowsAdmin 检测当前进程是否以管理员身份运行（Windows平台专用）
func isWindowsAdmin() (bool, error) {
	// 获取当前进程句柄
	hProcess, err := syscall.GetCurrentProcess()
	if err != nil {
		return false, err
	}

	// 打开当前进程的 token
	var hToken syscall.Token
	err = syscall.OpenProcessToken(hProcess, syscall.TOKEN_QUERY, &hToken)
	if err != nil {
		return false, err
	}
	defer hToken.Close()

	// 查询 TokenElevation 信息
	var elevation uint32
	var outLen uint32
	r, _, e := syscall.NewLazyDLL("advapi32.dll").NewProc("GetTokenInformation").Call(
		uintptr(hToken),
		uintptr(20), // TokenElevation
		uintptr(unsafe.Pointer(&elevation)),
		unsafe.Sizeof(elevation),
		uintptr(unsafe.Pointer(&outLen)),
	)
	if r == 0 {
		return false, e
	}

	return elevation != 0, nil
}
