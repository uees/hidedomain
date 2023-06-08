import { Breadcrumb, theme } from "antd"
import { Content } from "antd/es/layout/layout"
import { Outlet } from "react-router-dom";
import { useStore } from "../../../hooks";
import { observer } from "mobx-react-lite";

const LayoutContent: React.FC = () => {
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const { site, user } = useStore();

    return (
        <Content style={{ padding: '0 50px' }}>
            {
                user.token && site.breadcrumb ?
                    <Breadcrumb style={{ margin: '16px 0' }} items={site.breadcrumb} /> :
                    <div style={{ padding: '20px' }} ></div>
            }
            <div className="site-layout-content" style={{ background: colorBgContainer }}>
                <Outlet />
            </div>
        </Content>
    )
}

export default observer(LayoutContent);
