// import axios from "axios";
import request from "../utils/request";

export function getIP() {
  //return axios.get("https://ipapi.co/json/")
  return request.get("/ip");
}
