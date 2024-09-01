import React from "react"
import { useTitle } from "../../hooks";

const Profile: React.FC = () => {

    useTitle('个人信息')

    return (
        <>
        <h3>个人信息</h3>
        <div>开发中...  by wan</div>
        </>
    )
}

export default Profile;
