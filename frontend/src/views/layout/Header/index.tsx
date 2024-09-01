import React, { useEffect } from 'react';
import { observer } from "mobx-react-lite";
import { Layout, Menu } from 'antd';
import { useNavigate, Link } from 'react-router-dom';
import { HomeOutlined, BulbOutlined, UserOutlined, SettingOutlined, GlobalOutlined } from '@ant-design/icons';
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
                    key: '/domains',
                    icon: <BulbOutlined />,
                },
                {
                    label: '代理管理',
                    key: '/proxies',
                    icon: <GlobalOutlined />,
                },
                {
                    label: '设置',
                    key: '/settings',
                    icon: <SettingOutlined />,
                },
                {
                    label: user.username,
                    key: '/profile#user',
                    icon: <UserOutlined />,
                    children: [
                        {
                            label: '个人信息',
                            key: '/profile',
                        },
                        {
                            label: '登出',
                            key: '/logout',
                        },
                    ]
                },
            ])
        }
    }, [site, user.token, user.username])

    const navigate = useNavigate();

    const handleClick = ({ key }: { key: string }) => {
        if (key === "/logout") {
            user.logout()
            site.setMenu([])
            return navigate("/", { replace: true })
        }
        return navigate(key)
    }

    return (
        <Header style={{ display: 'flex', alignItems: 'center' }}>
            <div className="demo-logo"><Link to="/">Private Cloud</Link></div>
            <Menu theme="dark"
                mode="horizontal"
                items={site.menuItems}
                onClick={handleClick}
                style={{ flex: 1, minWidth: 0 }} />
        </Header>
    )
}

export default observer(LayoutHeader);
