import React, { useEffect } from "react"
import { useStore, useTitle } from "../../hooks";

const Profile: React.FC = () => {

    useTitle('个人信息')
    const { site } = useStore();

    useEffect(() => {
        site.setBreadcrumb([{ title: '主页' }, { title: '个人信息' }]);
    })

    return (
        <div>Hello, Profile</div>
    )
}

export default Profile;
