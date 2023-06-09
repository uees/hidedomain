import React from 'react';
import { Navigate } from "react-router-dom";
import { useStore } from "../hooks";

export const CheckPermission: React.FC = () => {
    const { user } = useStore()

    if (user.token) {
        return null;
    }

    user.logout()
    return (
        <Navigate replace to='/login' />
    )
}
