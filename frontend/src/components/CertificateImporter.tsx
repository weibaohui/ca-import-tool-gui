import React, { useState } from "react";
import { Button, Form, Card, Alert, Progress, Space, Typography } from "antd";
import { ImportOutlined, ClearOutlined } from "@ant-design/icons";
import { ImportCertificate, SelectCertificateFile } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

const { Text } = Typography;

interface CertificateImporterProps {
  onImportComplete: (result: main.ImportResult) => void;
}

const CertificateImporter: React.FC<CertificateImporterProps> = ({ onImportComplete }) => {
  const [form] = Form.useForm();
  const [importing, setImporting] = useState(false);
  const [progress, setProgress] = useState(0);
  const [selectedFilePath, setSelectedFilePath] = useState<string | null>(null);

  const handleFileSelect = async () => {
    try {
      const selectedPath = await SelectCertificateFile();
      if (selectedPath) {
        setSelectedFilePath(selectedPath);
        form.setFieldsValue({ filePath: selectedPath });
      }
    } catch (error: any) {
      console.error("文件选择失败:", error);
    }
  };

  const handleImport = async (values: any) => {
    if (!values.filePath) {
      onImportComplete({
        success: false,
        message: "请先选择证书文件",
        log: ""
      });
      return;
    }

    setImporting(true);
    setProgress(10); // 开始导入过程

    try {
      // 直接使用文件路径执行导入
      setProgress(30);

      const params: main.ImportParams = {
        file_path: values.filePath
      };

      const result = await ImportCertificate(params);
      setProgress(100);

      // 延迟一下以显示100%进度
      setTimeout(() => {
        setImporting(false);
        onImportComplete(result);
      }, 300);
    } catch (error: any) {
      setImporting(false);
      setProgress(0);
      onImportComplete({
        success: false,
        message: `导入失败: ${error.message || "未知错误"}`,
        log: ""
      });
    }
  };

  const handleReset = () => {
    form.resetFields();
    setProgress(0);
    setSelectedFilePath(null);
  };

  return (
    <Card title="证书导入工具" bordered={false}>
      <Form
        form={form}
        layout="vertical"
        onFinish={handleImport}
        disabled={importing}
      >
        <Alert
          message="安全提示"
          description="在导入过程中，系统可能会弹出输入密码的对话框，请您放心输入系统管理员密码以完成证书导入操作。"
          type="info"
          showIcon
          style={{ marginBottom: 20 }}
        />
        <Form.Item
          label="选择证书文件"
          name="filePath"
          rules={[{ required: true, message: '请选择证书文件!' }]}
        >
          <div>
            <Button onClick={handleFileSelect}>选择证书文件</Button>
            {selectedFilePath && (
              <div style={{ marginTop: 8 }}>
                <strong>已选择:</strong> {selectedFilePath}
              </div>
            )}
          </div>
        </Form.Item>

        {importing && (
          <div style={{ marginBottom: 20 }}>
            <Progress percent={progress} status="active" />
          </div>
        )}

        <Form.Item>
          <Space>
            <Button
              type="primary"
              htmlType="submit"
              icon={<ImportOutlined />}
              loading={importing}
            >
              导入证书
            </Button>
            <Button
              htmlType="button"
              onClick={handleReset}
              icon={<ClearOutlined />}
            >
              重置
            </Button>
          </Space>
        </Form.Item>
      </Form>
    </Card>
  );
};

export default CertificateImporter;