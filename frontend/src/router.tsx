import { createBrowserRouter } from "react-router-dom"
import ErrorPage from "./components/error-page";
import HomeRoute from './views/home/route';
import Layout from "./views/layout";
import Login from "./views/login";

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
                    {
                        //path: "contacts/:contactId/edit",
                        //element: <EditContact />,
                        //loader: contactLoader,
                        //action: editAction,
                    },
                ],
            },
            {
                path: "/login",
                element: <Login />,
            }
        ],
    },
]);

export default router
