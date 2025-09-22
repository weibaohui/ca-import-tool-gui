package main

// ImportResult 导入结果
type ImportResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Log     string `json:"log"`
}

// ImportParams 导入参数
type ImportParams struct {
	FilePaths []string `json:"file_paths"`
}

// CertificateInfo 证书信息
type CertificateInfo struct {
	Alias     string `json:"alias"`
	Subject   string `json:"subject"`
	Issuer    string `json:"issuer"`
	ValidFrom string `json:"valid_from"`
	ValidTo   string `json:"valid_to"`
}
