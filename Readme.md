# Golang Clean Web API (Dockerize) with a full sample project (Car Sale project)

## System Design Diagram

<p align="center"><img src='/docs/files/system_diagram.png' alt='Golang Web API System Design Diagram' /></p>

## Database Design Diagram

<p align="center"><img src='/docs/files/db_diagram.png' alt='Golang Web API System Design Diagram' /></p>

## Give a Star! :star:

If you like this repo or found it helpful, please give it a star. Thanks!

## Used Tools

1. [Gin as web framework](https://github.com/gin-gonic/gin)
2. [JWT for authentication and authorization](https://github.com/golang-jwt/jwt)
3. [Redis for caching](https://github.com/redis/redis)
4. [Elasticsearch for logging database](https://github.com/elastic/elasticsearch)
5. [Beat for log shipping](https://github.com/elastic/beats)
6. [Kibana as log viewer](https://github.com/elastic/kibana)
7. [Postgresql as main database engine](https://github.com/postgres/postgres)
8. [PgAdmin as database management tool](https://github.com/pgadmin-org/pgadmin4)
9. [Prometheus for metric database](https://github.com/prometheus/prometheus)
10. [Grafana for metric dashboards](https://github.com/grafana/grafana)
11. [Validator for endpoint input Validation](https://github.com/go-playground/validator)
12. [Viper for configurations](https://github.com/spf13/viper)
13. [Zap for logging](https://github.com/uber-go/zap)
14. [Zerolog for logging](https://github.com/rs/zerolog)
15. [Gorm as ORM](https://github.com/go-gorm/gorm)
16. [Swagger for documentation](https://github.com/swaggo/swag)
17. Docker compose to run project with all dependencies in docker

## How to run

### Run on local system

#### Start dependencies on docker
```bash 
docker compose -f "docker/docker-compose.yml" up -d setup elasticsearch kibana filebeat postgres pgadmin redis prometheus node-exporter alertmanager grafana
```

#### Install swagger and run app
```bash
cd src
go install github.com/swaggo/swag/cmd/swag@latest
cd src/cmd
go run main.go
```

##### Address: [http://localhost:5005](http://localhost:5005)

#### Stop
```bash
docker compose -f "docker/docker-compose.yml" down
```

#### Examples

##### Login
```bash
curl -X 'POST' \
  'http://localhost:5005/api/v1/users/login-by-username' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "12345678",
  "username": "admin"
}'
```

##### Sample filters request body

###### City filter and sort

```json
{
  "filter": {
    "Name": {
      "filterType": "text",
      "from": "t", 
      "type": "contains"
    } 
  },
  "pageNumber": 1,
  "pageSize": 10,
  "sort": [
    {
      "colId": "name",
      "sort": "desc"
    }
  ]
}
```


###### City in range filter 

```json
{
  "filter": {
    "Id": { // Column name
      "filterType": "number", // number, text,...
      "from": "1", 
      "to": "7", 
      "type": "inRange" // contains, equals,...
    } 
  },
  "pageNumber": 1,
  "pageSize": 10,
  "sort": [
    {
      "colId": "name",
      "sort": "desc"
    }
  ]
}
```

### Run project with dependencies on Docker

```bash
docker compose -f "docker/docker-compose.yml" up -d --build
```

#### Web API  Run in docker  [http://localhost:9001](http://localhost:9001)

```
Token Url: http://localhost:9001/api/v1/users/login-by-username
Username: admin
Password: 12345678
```

#### Kibana  [http://localhost:5601](http://localhost:5601)

```
Username: elastic
Password: @aA123456
```

#### Prometheus  [http://localhost:9090](http://localhost:9090)

#### Grafana  [http://localhost:3000](http://localhost:3000)

```
Username: admin
Password: foobar
```

#### PgAdmin  [http://localhost:8090](http://localhost:8090)

```
Username: h.naimaei@gmail.com
Password: 123456
```

Postgres Server info:

```
Host: postgres_container
Port: 5432
Username: postgres
Password: admin
```

### Docker Stop

```bash
docker compose -f 'docker/docker-compose.yml' --project-name 'docker' down
```

### Linux

0. build Project and copy configuration

```bash
/src > go build -o ../prod/server ./cmd/main.go
/src > mkdir ../prod/config/ && cp config/config-production.yml ../prod/config/config-production.yml
```

1. Create systemd unit

```bash
sudo vi /lib/systemd/system/go-api.service
```

2. Service config

```
[Unit]
Description=go-api

[Service]
Type=simple
Restart=always
RestartSec=20s
ExecStart=/home/hamed/github/golang-clean-web-api/prod/server
Environment="APP_ENV=production"
WorkingDirectory=/home/hamed/github/golang-clean-web-api/prod
[Install]
WantedBy=multi-user.target
```

3. Start service

```bash
sudo systemctl start go-api
```

4. Stop service

```bash
sudo systemctl stop go-api
```

5. Show service logs

```bash
sudo journalctl -u go-api -e
```

## Project preview

## Swagger

<p align="center"><img src='/docs/files/swagger.png' alt='Golang Web API preview' /></p>

## Grafana

<p align="center"><img src='/docs/files/grafana.png' alt='Golang Web API grafana dashboard' /></p>

## Kibana

<p align="center"><img src='/docs/files/kibana.png' alt='Golang Web API grafana dashboard' /></p>
