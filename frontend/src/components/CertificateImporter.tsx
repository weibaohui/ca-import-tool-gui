import React, { useState } from "react";
import { Button, Form, Card, Alert, Progress, Space, Typography, List } from "antd";
import { ImportOutlined, ClearOutlined } from "@ant-design/icons";
import { ImportCertificate, SelectCertificateFiles } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

const { Text } = Typography;

interface CertificateImporterProps {
  onImportComplete: (result: main.ImportResult) => void;
}

const CertificateImporter: React.FC<CertificateImporterProps> = ({ onImportComplete }) => {
  const [form] = Form.useForm();
  const [importing, setImporting] = useState(false);
  const [progress, setProgress] = useState(0);
  const [selectedFilePaths, setSelectedFilePaths] = useState<string[]>([]);

  const handleFileSelect = async () => {
    try {
      const selectedPaths = await SelectCertificateFiles();
      if (selectedPaths && selectedPaths.length > 0) {
        setSelectedFilePaths(selectedPaths);
        form.setFieldsValue({ filePaths: selectedPaths });
      }
    } catch (error: any) {
      console.error("文件选择失败:", error);
    }
  };

  const handleImport = async (values: any) => {
    if (!values.filePaths || values.filePaths.length === 0) {
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
        file_paths: values.filePaths
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
    setSelectedFilePaths([]);
  };

  return (
    <Card title="证书导入工具" bordered={false}>
      <Form
        form={form}
        layout="vertical"
        onFinish={handleImport}
        disabled={importing}
      >
        {selectedFilePaths.length == 1 && (
          <Alert
            description="在导入过程中，系统可能会弹出输入密码的对话框，请您放心输入系统管理员密码以完成证书导入操作。"
            type="info"
            style={{ marginBottom: 20 }}
          />
        )}


        {selectedFilePaths.length > 1 && (
          <Alert
            description="您选择了多个证书文件进行导入，系统可能会为每个证书文件分别弹出输入密码的对话框，请您耐心输入。"
            type="warning"
            style={{ marginBottom: 20 }}
          />
        )}
        <Form.Item
          label="选择证书文件"
          name="filePaths"
          rules={[{ required: true, message: '请选择证书文件!' }]}
        >
          <div>
            <Button onClick={handleFileSelect}>选择证书文件</Button>
            {selectedFilePaths.length > 0 && (
              <div style={{ marginTop: 8 }}>
                <strong>已选择 {selectedFilePaths.length} 个文件:</strong>
                <List
                  size="small"
                  bordered
                  dataSource={selectedFilePaths}
                  renderItem={item => <List.Item>{item}</List.Item>}
                  style={{ maxHeight: 200, overflow: 'auto', marginTop: 8 }}
                />
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