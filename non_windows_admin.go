//go:build !windows

package main

// isWindowsAdmin 非Windows平台的管理员权限检测（始终返回true）
func isWindowsAdmin() (bool, error) {
	return true, nil
}
