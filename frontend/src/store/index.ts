import React from "react";
import UserStore from "./user";
import SiteStore from "./site";

const userStore = new UserStore();
const siteStore = new SiteStore();

export const store = {
  userStore,
  siteStore,
};

export const storeContext = React.createContext({
  user: userStore,
  site: siteStore,
});
