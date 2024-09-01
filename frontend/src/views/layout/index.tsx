import { Layout as AntdLayout } from "antd"
import { Link, useLoaderData } from "react-router-dom";
import { Footer } from "antd/es/layout/layout"
import LayoutContent from "./Content";
import LayoutHeader from "./Header";
import { CheckPermission } from "../permission";
import { store } from '../../store'
import UserStore from "../../store/user";
import '../../styles/style.css';

export async function loader() {
    const { userStore } = store;
    if (userStore.token && !userStore.username) {
        await userStore.loadInfo();
    }

    return { user: userStore }
}

const Layout: React.FC = () => {

    const today = new Date()
    const year = today.getFullYear()

    const { user }= useLoaderData() as {user: UserStore};

    return (
        <>
            <CheckPermission />
            <AntdLayout className="layout">
                <LayoutHeader />
                <LayoutContent />
                <Footer style={{ textAlign: 'center' }}>
                    Zues.PUB ©2017-{year} {!user.token && <Link to="/login">Login</Link>}
                </Footer>
            </AntdLayout>
        </>
    )
}

export default Layout;
