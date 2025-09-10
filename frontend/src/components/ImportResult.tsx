import React from "react";
import { Card, Alert, Typography, Divider } from "antd";
import { main } from "../../wailsjs/go/models";

const { Title, Text } = Typography;

interface ImportResultProps {
  result: main.ImportResult;
}

const ImportResult: React.FC<ImportResultProps> = ({ result }) => {
  return (
    <Card title="导入结果" bordered={false}>
      <Alert
        message={result.message}
        type={result.success ? "success" : "error"}
        showIcon
      />
      
      {result.log && (
        <>
          <Divider orientation="left">详细日志</Divider>
          <div style={{ 
            backgroundColor: '#f6f6f6', 
            padding: '12px', 
            borderRadius: '4px',
            maxHeight: '200px',
            overflow: 'auto',
            whiteSpace: 'pre-wrap'
          }}>
            <Text>{result.log}</Text>
          </div>
        </>
      )}
    </Card>
  );
};

export default ImportResult;