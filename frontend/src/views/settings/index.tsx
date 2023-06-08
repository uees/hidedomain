import React, { useEffect } from "react"
import { useStore, useTitle } from "../../hooks";

const Settings: React.FC = () => {

    useTitle('设置')
    const { site } = useStore();

    useEffect(() => {
        site.setBreadcrumb([{ title: '主页' }, { title: '设置' }]);
    })

    return (
        <div>Hello, Settings</div>
    )
}

export default Settings;
