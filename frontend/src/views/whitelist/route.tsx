import View, { loader, deleteRuleAction } from "."

const route = {
    path: "/domains/:domain/whitelist",
    element: <View />,
    loader: loader,
}

export default route;

export const deleteRuleRoute = {
    path: "/domains/:domain/whitelist/:ruleid/destroy",
    errorElement: <div>Oops! There was an error.</div>,
    action: deleteRuleAction,
}
