import { IWhitelist } from "../interfaces/models";
import request from "../utils/request";

export function showList(domain: string) {
  return request.get(`/domains/${domain}/whitelist`);
}

export function clearList(domain: string) {
  return request.delete(`/domains/${domain}/whitelist`);
}

export const addIPRule = (domain: string, rule: IWhitelist) => {
  return request.post(`/domains/${domain}/whitelist`, rule);
};

export const showIPRule = (domain: string, ruleID: number | string) => {
  return request.get(`/domains/${domain}/whitelist/${ruleID}`);
};

export const deleteIPRule = (domain: string, ruleID: number | string) => {
  return request.delete(`/domains/${domain}/whitelist/${ruleID}`);
};

export const updateIPRule = (
  domain: string,
  ruleID: number | string,
  rule: IWhitelist
) => {
  return request.patch(`/domains/${domain}/whitelist/${ruleID}`, rule);
};
