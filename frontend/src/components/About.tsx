import React from "react";
import { Descriptions, Typography, List, Card } from "antd";

const { Title, Paragraph } = Typography;

const About: React.FC = () => {
    const supportedSystems = [
        {
            name: "Windows",
            versions: ["Windows 10", "Windows 11"],
            description: "支持通过certutil命令导入证书"
        },
        {
            name: "macOS",
            versions: ["macOS 10.15 (Catalina)及以上版本"],
            description: "支持通过security命令导入证书"
        },
        {
            name: "Linux",
            versions: [
                "Ubuntu 18.04 LTS及以上版本",
                "Debian 10及以上版本",
                "CentOS 7及以上版本",
                "Red Hat Enterprise Linux 7及以上版本",
                "Fedora 32及以上版本",
                "openSUSE Leap 15及以上版本",
                "Arch Linux (最新版本)",
                "Alpine Linux 3.12及以上版本"
            ],
            description: "支持update-ca-certificates和update-ca-trust命令"
        }
    ];

    return (
        <div style={{ padding: "20px" }}>
            <Title level={3}>关于CA证书导入工具</Title>
            <Paragraph>
                CA证书导入工具是一个跨平台的桌面应用程序，旨在简化在不同操作系统上导入和管理CA证书的过程。
                该工具支持Windows、macOS和主要的Linux发行版。
            </Paragraph>

            <Title level={4} style={{ marginTop: "30px" }}>支持的操作系统</Title>
            <List
                grid={{ gutter: 16, column: 1 }}
                dataSource={supportedSystems}
                renderItem={system => (
                    <List.Item>
                        <Card title={system.name} style={{ width: "100%" }}>
                            <Descriptions column={1} layout="vertical">
                                <Descriptions.Item label="支持版本">
                                    <List
                                        size="small"
                                        dataSource={system.versions}
                                        renderItem={version => <List.Item>{version}</List.Item>}
                                    />
                                </Descriptions.Item>
                                <Descriptions.Item label="实现方式">
                                    {system.description}
                                </Descriptions.Item>
                            </Descriptions>
                        </Card>
                    </List.Item>
                )}
            />

            <Title level={4} style={{ marginTop: "30px" }}>技术信息</Title>
            <Descriptions column={1} bordered>
                <Descriptions.Item label="开发框架">Wails (Go + React)</Descriptions.Item>
                <Descriptions.Item label="前端技术">React + TypeScript + Ant Design</Descriptions.Item>
                <Descriptions.Item label="后端技术">Go语言</Descriptions.Item>
                <Descriptions.Item label="构建工具">Vite</Descriptions.Item>
            </Descriptions>

            <Title level={4} style={{ marginTop: "30px" }}>许可证</Title>
            <Paragraph>
                本软件是免费软件，遵循MIT许可证发布。
            </Paragraph>
        </div>
    );
};

export default About;