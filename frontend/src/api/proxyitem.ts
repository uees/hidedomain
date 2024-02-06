import { IProxyItem } from "../interfaces/models";
import request from "../utils/request";

export function proxiesList() {
  return request.get("/proxies");
}

export function getProxyItem(id: string) {
  return request.get(`/proxies/${id}`);
}

export function createProxyItem(data: IProxyItem) {
  return request.post("/proxies", data);
}

export const updateProxyItem = (id: string, data: IProxyItem) => {
  return request.patch(`/proxies/${id}`, data);
};

export const deleteProxyItem = (id: string) => {
  return request.delete(`/proxies/${id}`);
};
