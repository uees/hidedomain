import { useTitle } from "../../hooks";
import privateCloudImage from '../../img/PrivateCloud.png';

const Home: React.FC = () => {

    useTitle('Welcome to Private Cloud!')

    return (
        <>
        <h1>Welcome to Private Cloud!</h1>
        <div>
        <p><b>What is a private cloud?</b></p>
        <p>A private cloud is a form of cloud computing service built for the sole use of an organization. A private cloud
        provides effective control over data, security, and quality of service. The organization owns the infrastructure
        and controls the deployment of its own network and application services on top of it. A private cloud can be
        built by the organization's own ICT department, or by a dedicated private cloud provider. The owner of the
        private cloud does not share resources with other enterprises or organizations, and the core attribute of the
        private cloud is resource exclusiveness.</p>
        <p><b>Characteristics and Value of Private Cloud</b></p>
        <p>There are three types of modes for enterprises to deploy cloud computing services: public cloud (Public Cloud), private cloud (Private Cloud), and hybrid cloud (Hybird Cloud). Among them, a private cloud is a proprietary cloud computing system built internally by a cloud computing service provider for a specific organization. Private cloud systems exist within the firewall of the organization's data center, or are deployed in a secure colocation facility, and serve only that organization.</p>
        <p>The characteristics and value of private cloud are:</p>
        <p>Security: The private cloud is only available to specific users, not the general public, and is deployed behind the organization's own firewall, thus providing a higher level of security and privacy, ensuring that sensitive data cannot be accessed by third parties.</p>
        <p>Stable SLA: The private cloud is usually deployed in the organization's own data center. When the organization internally accesses the resources in the private cloud, the SLA is very stable.</p>
        <p>Self-controllable: Private cloud customers are free to purchase their favorite hardware and software instead of those provided by public cloud service providers.</p>
        <p>Strong customization: private cloud customers can customize computing, storage and network in any way they want, and run their own customized software and management platform. Private cloud is customized and built around the actual business needs of the organization.</p>
        <p><b>Private Cloud Architecture</b></p>
        <img src={privateCloudImage} alt="Private Cloud Architecture" />
        </div>
        <p><em>Thank you for using Private Cloud.</em></p>
        </>
    )
}

export default Home;
