import React from 'react';
import { Navigate } from "react-router-dom";
import useStore from "../hooks/useStore";

export const CheckPermission: React.FC = () => {
    const { user } = useStore()

    if (user.token) {
        if (!user.username) {
            user.loadInfo();
        }
        return null;
    }

    user.logout()
    return (
        <Navigate replace to='/login' />
    )
}
