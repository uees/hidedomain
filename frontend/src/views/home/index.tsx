import { useEffect } from "react";
import { useStore, useTitle } from "../../hooks";

const Home: React.FC = () => {

    useTitle('Preventive active detection')

    const { site } = useStore();

    useEffect(() => {
        site.setBreadcrumb([{ title: '主页' }]);
    })

    return (
        <div>Hello. it is essential to prevent active detection.</div>
    )
}

export default Home;
