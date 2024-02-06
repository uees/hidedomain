import { makeAutoObservable, runInAction } from "mobx";
import { getToken, setToken, removeToken } from "../utils/auth";
import { login, profile } from "../api/auth";
import { IUserForm } from "../interfaces";
import { ITokenData, IUser } from "../interfaces/models";

export default class UserStore {
  username = "";
  email = "";
  email_verified = false;
  role = "";
  login_ip = "";
  login_ua = "";
  login_at = "";
  token = getToken();
  constructor() {
    makeAutoObservable(this);
  }

  async login(user: IUserForm) {
    const { data } = await login(user.username.trim(), user.password);
    // console.log(data)
    const tokenData = data.data as ITokenData;
    runInAction(() => {
      this.token = tokenData.token;
      setToken(tokenData.token, tokenData.refresh_expires_in);
    });
  }

  async loadInfo() {
    const { data } = await profile();
    const userData = data.data as IUser;

    runInAction(() => {
      this.email = userData.email;
      this.username = userData.username;
      this.email_verified = userData.email_verified;
      this.login_at = userData.login_at;
      this.login_ip = userData.login_ip;
      this.login_ua = userData.login_ua;
      this.role = userData.role;
    });
  }

  logout() {
    setToken("", 0);
    removeToken();
    this.reset();
  }

  reset() {
    this.email = "";
    this.username = "";
    this.email_verified = false;
    this.login_at = "";
    this.login_ip = "";
    this.login_ua = "";
    this.role = "";
    this.token = "";
  }
}
