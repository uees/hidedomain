# hidedomain

### POST /api/token

params:
    username: string
    password: string

### GET /api/profile

### GET /api/domains

### POST /api/domains

params:
    domain: string
    zone_id: string
    account_id: string
    api_key: string
    mode: string (local or cf)

### PATCH /api/domains/:domain
    domain: string
    mode: string (local or cf)

### GET /api/domains/:domain

### DELETE /api/domains/:domain

### POST /api/domains/:domain/whitelist

params:
    ip: string

### GET /api/domains/:domain/whitelist

### DELETE /api/domains/:domain/whitelist

### DELETE /api/domains/:domain/whitelist/:ruleid

### PATCH /api/domains/:domain/whitelist/:ruleid
