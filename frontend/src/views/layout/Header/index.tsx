import React from 'react';
import { Layout, Menu } from 'antd';
import { HomeOutlined, BulbOutlined, FileOutlined, ProfileOutlined } from '@ant-design/icons';
import useStore from '../../../hooks/useStore';

const { Header } = Layout;

const LayoutHeader: React.FC = () => {

    const { site, user } = useStore()

    if (user.token) {
        site.setMenu([
            {
                label: '主页',
                key: 'home',
                icon: <HomeOutlined />,
            },
            {
                label: '域名管理',
                key: 'domain',
                icon: <BulbOutlined />,
            },
            {
                label: '白名单',
                key: 'whitelist',
                icon: <FileOutlined />,
            },
            {
                label: 'Profile',
                key: 'profile',
                icon: <ProfileOutlined />,
            },
        ])
    }

    return (
        <Header>
            <div className="demo-logo" />
            {user.username && <Menu theme="dark" mode="horizontal" items={site.menuItems} />}
        </Header>
    )
}

export default LayoutHeader;
