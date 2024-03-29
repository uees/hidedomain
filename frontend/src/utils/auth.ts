import Cookies from "js-cookie";

const TokenKey = "ZUES_TOKEN";

export function getToken() {
  return Cookies.get(TokenKey);
}

export function setToken(token: string, expires_in: number) {
  const options = {
    expires: 14, // default 2 weeks
  };
  if (expires_in) {
    options.expires = expires_in / (60 * 60 * 24); // 秒转天
  }
  return Cookies.set(TokenKey, token, options);
}

export function removeToken() {
  return Cookies.remove(TokenKey);
}
