# Active Charity Backend

### GRPC-PORT: 6500
### HTTP-PORT: 6501
### POSTGRESQL-PORT: 5432

## LOCAL CONFIGURATION
```yaml
server:
  ip:   "0.0.0.0"
  port: "65000"

gateway:
  ip: "0.0.0.0"
  port: "65001"

db:
  host:     "postgresql"
  port:     "5432"
  user:     "postgres"
  password: "12345678"
  db_name:  "active_charity"
  ssl_mode: "disable"

jwt:
  secret_key:      "1234"
  expiration_time: "168h"
```

### REQUIREMENTS
- go 1.21.1
- docker & docker-compose
- postgresql

## DOCKER

### Build
```shell
docker-compose build
```
### Run
```shell
docker-compose up
```