import { createBrowserRouter } from "react-router-dom"
import Layout, { loader } from "./views/layout";
import ErrorPage from "./views/error-page";
import HomeRoute from './views/home/route';
import LoginRoute from "./views/login/route";
import DomainRoute, {
    addDomainRoute,
    deleteDomainRoute,
    editDomainRoute,
} from "./views/domain/route";
import ProxyItemRoute, {
    addProxyItemRoute,
    deleteProxyItemRoute,
    editProxyItemRoute,
} from "./views/proxy/route";
import ProfileRoute from "./views/profile/route";
import SettingsRoute from "./views/settings/route";
import WhitelistRoute, { deleteRuleRoute } from "./views/whitelist/route";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Layout />,
        loader: loader,
        errorElement: <ErrorPage />,
        children: [
            { ...HomeRoute },
            { ...DomainRoute },
            { ...addDomainRoute },
            { ...editDomainRoute },
            { ...deleteDomainRoute },
            { ...ProxyItemRoute },
            { ...addProxyItemRoute },
            { ...deleteProxyItemRoute },
            { ...editProxyItemRoute },
            { ...deleteRuleRoute },
            { ...WhitelistRoute },
            { ...SettingsRoute },
            { ...ProfileRoute },
            { ...LoginRoute },
        ],
    },
]);

export default router
