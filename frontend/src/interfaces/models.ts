export interface IUser {
    username: string;
    email: string;
    email_verified: boolean;
    role: string;
    login_ip: string;
    login_ua: string;
    login_at: string;
}

export interface IOption {
    name: string;
    value: string;
    memo: string;
}

export interface IDomain {
    name: string;
    mode: string;
    memo: string;
}

export interface IWhitelist {
    id: string;
    domain: IDomain;
    ip: string;
    memo: string;
}

export interface ITokenData {
    token: string;
    refresh_expires_in: number;
}
