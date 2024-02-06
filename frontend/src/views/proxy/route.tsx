import View, { loader } from "."
import ProxyItemForm, { deleteProxyItemAction, loader as formLoader } from './form'

const route = {
    path: "/proxies",
    element: <View />,
    loader: loader,
}

export default route;

export const addProxyItemRoute = {
    path: "/proxies/add",
    loader: formLoader,
    element: < ProxyItemForm />,
}

export const editProxyItemRoute = {
    path: "/proxies/:id/edit",
    loader: formLoader,
    element: <ProxyItemForm />,
}

export const deleteProxyItemRoute = {
    path: "/proxies/:id/destroy",
    action: deleteProxyItemAction,
    errorElement: <div>Oops! There was an error.</div>,
}
