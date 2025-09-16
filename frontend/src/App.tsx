import React, { useState } from "react";
import { Layout, Typography, Card, Tabs } from "antd";
import logo from './assets/images/logo-universal.png';
import CertificateImporter from "./components/CertificateImporter";
import ImportResultComponent from "./components/ImportResult";
import SystemInfo from "./components/SystemInfo";
import { main } from "../wailsjs/go/models";

const { Header, Content, Footer } = Layout;
const { Title, Text } = Typography;
const { TabPane } = Tabs;

function App() {
    const [importResult, setImportResult] = useState<main.ImportResult | null>(null);
    const [activeTab, setActiveTab] = useState("import");

    const handleImportComplete = (result: main.ImportResult) => {
        setImportResult(result);
        setActiveTab("result");
    };

    return (
        <Layout style={{ minHeight: "100vh" }}>
            <Header style={{ color: "#fff" }}>
                <Title level={3} style={{ color: "white", margin: 0 }}>
                    🚀 CA证书导入工具
                </Title>
            </Header>
            <Content style={{ padding: "20px" }}>
                <Card style={{ maxWidth: 800, margin: "auto" }}>
                    <Tabs
                        activeKey={activeTab}
                        onChange={setActiveTab}
                        items={[
                            {
                                key: "import",
                                label: "证书导入",
                                children: (
                                    <CertificateImporter onImportComplete={handleImportComplete} />
                                ),
                            },
                            {
                                key: "result",
                                label: "导入结果",
                                children: importResult ? (
                                    <ImportResultComponent result={importResult} />
                                ) : (
                                    <div>请先执行证书导入操作</div>
                                ),
                                disabled: !importResult,
                            },
                            {
                                key: "system",
                                label: "系统信息",
                                children: <SystemInfo />,
                            },
                        ]}
                    />
                </Card>
            </Content>
            <Footer style={{ textAlign: "center" }}>
                CA证书导入工具 ©2025
            </Footer>
        </Layout>
    );
}

export default App;