import { IDomain } from '../interfaces/models'
import request from '../utils/request'

export function domainList() {
    return request.get('/domains')
}

export function createDomain(domain: IDomain) {
    return request.post('/domains', domain)
}

export const updateDomain = (name: string, domain: IDomain) => {
    return request.patch(`/domains/${name}`, domain)
}

export const deleteDomain = (name: string) => {
    return request.delete(`/domains/${name}`)
}
