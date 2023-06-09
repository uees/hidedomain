import View, { loader } from "."
import DomainForm, { deleteDomainAction, loader as formLoader } from './form'

const route = {
    path: "/domains",
    element: <View />,
    loader: loader,
}

export default route;

export const addDomainRoute = {
    path: "/domains/add",
    loader: formLoader,
    element: < DomainForm />,
}

export const editDomainRoute = {
    path: "/domains/:domain/edit",
    loader: formLoader,
    element: <DomainForm />,
}

export const deleteDomainRoute = {
    path: "/domains/:domain/destroy",
    action: deleteDomainAction,
    errorElement: <div>Oops! There was an error.</div>,
}
