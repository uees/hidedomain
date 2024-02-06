import request from "../utils/request";

export function login(username: string, password: string) {
  const data = {
    username,
    password,
  };
  return request({
    url: "/token",
    method: "post",
    data,
  });
}

export function profile() {
  return request.get("/profile");
}
