import { Link } from "react-router-dom"

const Navbar = (props: any) => {
    return (
        <div>
            <nav>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/control-panel">Control Panel</Link></li>
                </ul>
            </nav>
        </div>
    )
}
