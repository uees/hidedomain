import React from "react"
import { useParams } from "react-router-dom";

const Whitelist: React.FC = () => {

    let { domain } = useParams();

    return (
        <div>Hello, There are <b>{domain}</b> 's Whitelist</div>
    )
}

export default Whitelist;
