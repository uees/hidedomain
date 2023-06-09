import { IWhitelist } from '../interfaces/models'
import request from '../utils/request'

export function showList(domain: string) {
    return request.get(`/domains/${domain}/whitelist`)
}

export function clearList(domain: string) {
    return request.delete(`/domains/${domain}/whitelist`)
}

export const addIPRule = (domain: string, rule: IWhitelist) => {
    return request.post(`/domains/${domain}/whitelist`, rule)
}

export const showIPRule = (ruleID: number | string) => {
    return request.get(`whitelist/${ruleID}`);
}

export const deleteIPRule = (ruleID: number | string) => {
    return request.delete(`/whitelist/${ruleID}`)
}

export const updateIPRule = (ruleID: number | string, rule: IWhitelist) => {
    return request.patch(`/whitelist/${ruleID}`, rule)
}
