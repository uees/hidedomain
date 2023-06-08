import useStore from "../../hooks/useStore";

const Home: React.FC = () => {
    const { site } = useStore();
    site.setPageTitle('Preventive active detection');
    site.setBreadcrumb([{ title: 'Home' }]);

    return (
        <div>Hello. it is essential to prevent active detection.</div>
    )
}

export default Home;
