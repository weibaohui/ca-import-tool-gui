import React, { useEffect, useState } from "react";
import { Card, Descriptions, Spin, Alert } from "antd";
import { InfoCircleOutlined } from "@ant-design/icons";
import { GetSystemInfo } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

const SystemInfo: React.FC = () => {
    const [systemInfo, setSystemInfo] = useState<main.SystemInfo | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchSystemInfo = async () => {
            try {
                setLoading(true);
                const info = await GetSystemInfo();
                setSystemInfo(info);
            } catch (err: any) {
                setError("获取系统信息失败: " + (err.message || "未知错误"));
            } finally {
                setLoading(false);
            }
        };

        fetchSystemInfo();
    }, []);

    // 将GOOS值转换为可读的系统名称
    const getOSName = (os: string) => {
        switch (os) {
            case "darwin":
                return "macOS";
            case "windows":
                return "Windows";
            case "linux":
                return "Linux";
            default:
                return os;
        }
    };

    return (
        <Card
            title={
                <div style={{ display: "flex", alignItems: "center", gap: 8 }}>
                    <InfoCircleOutlined />
                    <span>系统信息</span>
                </div>
            }
            bordered={false}
        >
            {error && (
                <Alert
                    message="错误"
                    description={error}
                    type="error"
                    showIcon
                    style={{ marginBottom: 16 }}
                />
            )}

            {loading ? (
                <div style={{ textAlign: "center", padding: "20px" }}>
                    <Spin size="large" />
                    <div style={{ marginTop: 10 }}>正在加载系统信息...</div>
                </div>
            ) : systemInfo ? (
                <Descriptions bordered column={1}>
                    <Descriptions.Item label="操作系统">
                        {getOSName(systemInfo.os)} ({systemInfo.arch})
                    </Descriptions.Item>
                    <Descriptions.Item label="Go版本">
                        {systemInfo.go_version}
                    </Descriptions.Item>
                    <Descriptions.Item label="应用名称">
                        {systemInfo.app_name}
                    </Descriptions.Item>
                    <Descriptions.Item label="应用版本">
                        {systemInfo.app_ver}
                    </Descriptions.Item>
                </Descriptions>
            ) : (
                !error && (
                    <Alert
                        message="无数据"
                        description="暂无系统信息可显示"
                        type="info"
                        showIcon
                    />
                )
            )}
        </Card>
    );
};

export default SystemInfo;