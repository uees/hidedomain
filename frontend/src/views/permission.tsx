import React from 'react';
import { Navigate } from "react-router-dom";
import { getToken } from "../utils/auth";
import useStore from "../hooks/useStore";

export const CheckPermission: React.FC = () => {
    const hasToken = getToken()
    const { user } = useStore()

    if (hasToken) {
        return null;
    }

    user.logout()
    return (
        <Navigate replace to='/login' />
    )
}
