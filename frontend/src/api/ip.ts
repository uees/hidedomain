import axios from "axios";

export function getIP() {
    return axios.get("https://ipapi.co/json/")
}
