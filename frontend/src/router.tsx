import { createBrowserRouter } from "react-router-dom"
import Layout from "./views/layout";
import ErrorPage from "./views/error-page";
import HomeRoute from './views/home/route';
import LoginRoute from "./views/login/route";
import DomainRoute from "./views/domain/route";
import ProfileRoute from "./views/profile/route";
import SettingsRoute from "./views/settings/route";
import WhitelistRoute from "./views/whitelist/route";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Layout />,
        errorElement: <ErrorPage />,
        children: [
            {
                errorElement: <ErrorPage />,
                children: [
                    { ...HomeRoute },
                    { ...DomainRoute },
                    { ...WhitelistRoute },
                    { ...SettingsRoute },
                    { ...ProfileRoute },
                ],
            },
            { ...LoginRoute },
        ],
    },
]);

export default router
