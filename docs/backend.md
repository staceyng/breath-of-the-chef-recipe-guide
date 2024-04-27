# Backend Documentation
## Functionality
- Search for recipes by ingredients

## Backend
### Database Design

Technology
- Language: Golang
- Database: Postgres

Libraries
- [pq](https://github.com/lib/pq): postgres driver for golang
- [gin](https://gin-gonic.com/docs/): http web framework for golang

### Setup
Database setup
1. psql
2. create database botc
3. cd backend/db/migrations run `GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://<user>:<password>@localhost:5432/botc" goose up/down
