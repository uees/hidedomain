import { Layout as AntdLayout } from "antd"
import { Footer } from "antd/es/layout/layout"
import LayoutContent from "./Content";
import LayoutHeader from "./Header";
import { CheckPermission } from "../permission";
import '../../styles/style.css';

const Layout: React.FC = () => {

    return (
        <>
            <CheckPermission />
            <AntdLayout className="layout">
                <LayoutHeader />
                <LayoutContent />
                <Footer style={{ textAlign: 'center' }}>©2023 Zues.PUB</Footer>
            </AntdLayout>
        </>
    )
}

export default Layout;
