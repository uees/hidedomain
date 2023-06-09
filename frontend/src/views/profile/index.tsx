import React from "react"
import { useTitle } from "../../hooks";

const Profile: React.FC = () => {

    useTitle('个人信息')

    return (
        <div>Hello, Profile</div>
    )
}

export default Profile;
