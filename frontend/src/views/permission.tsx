import React from 'react';
import { Navigate, useLocation } from "react-router-dom";
import { useStore } from "../hooks";

export const CheckPermission: React.FC = () => {
    const { user } = useStore()
    const location = useLocation()

    if (!user.token && location.pathname !== "/") {
        user.logout()
        return (
            <Navigate replace to='/login' />
        )
    }

    return null
}
