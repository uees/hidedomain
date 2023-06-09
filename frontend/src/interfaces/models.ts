export interface IUser {
    id?: string | number;
    username: string;
    email: string;
    email_verified: boolean;
    role: string;
    login_ip: string;
    login_ua: string;
    login_at: string;
}

export interface IOption {
    id?: string | number;
    name: string;
    value: string;
    memo: string;
}

export interface IDomain {
    id?: string | number;
    key?: string | number;
    name: string;
    mode: string;
    zone_id?: string;
    account_id?: string;
    api_key?: string;
    memo?: string;
}

export interface IWhitelist {
    id?: string | number;
    domain?: IDomain;
    ip: string;
    memo: string;
}

export interface ITokenData {
    token: string;
    refresh_expires_in: number;
}
