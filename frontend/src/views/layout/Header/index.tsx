import React, { useEffect } from 'react';
import { observer } from "mobx-react-lite";
import { Layout, Menu } from 'antd';
import { useNavigate } from 'react-router-dom';
import { HomeOutlined, BulbOutlined, UserOutlined, SettingOutlined } from '@ant-design/icons';
import { useStore } from '../../../hooks';

const { Header } = Layout;

const LayoutHeader: React.FC = () => {

    const { site, user } = useStore()

    useEffect(() => {
        if (user.token) {
            site.setMenu([
                {
                    label: '主页',
                    key: '/',
                    icon: <HomeOutlined />,
                },
                {
                    label: '域名管理',
                    key: '/domain',
                    icon: <BulbOutlined />,
                },
                {
                    label: '设置',
                    key: '/settings',
                    icon: <SettingOutlined />,
                },
                {
                    label: '个人设置',
                    key: '/profile#user',
                    icon: <UserOutlined />,
                    children: [
                        {
                            label: 'Profile',
                            key: '/profile',
                        },
                        {
                            label: '退出系统',
                            key: '/logout',
                        },
                    ]
                },
            ])
        }
    }, [site, user.token])

    const navigate = useNavigate();

    const handleClick = ({ key }: { key: string }) => {
        if (key === "/logout") {
            user.logout()
            return navigate("/login", { replace: true })
        }
        return navigate(key)
    }

    return (
        <Header>
            <div className="demo-logo" />
            <Menu theme="dark" mode="horizontal" items={site.menuItems} onClick={handleClick} />
        </Header>
    )
}

export default observer(LayoutHeader);
