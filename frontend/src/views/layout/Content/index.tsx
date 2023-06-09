import { theme } from "antd"
import { Content } from "antd/es/layout/layout"
import { Outlet } from "react-router-dom";

const LayoutContent: React.FC = () => {
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    return (
        <Content style={{ padding: '0 50px' }}>
            <div className="site-layout-content"
                style={{ background: colorBgContainer, padding: '20px 8px 8px', marginTop: '20px' }}
            >
                <Outlet />
            </div>
        </Content>
    )
}

export default LayoutContent;
