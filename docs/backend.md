# Backend Documentation
## Functionality
1. Search a recipe by recipe name
2. Show all recipes for a category
3. Search for recipes by ingredients

## Backend
### Database Design

Technology
- Language: Golang
- Database: Postgres

Libraries
- [pq](https://github.com/lib/pq): postgres driver for golang
- [gin](https://gin-gonic.com/docs/): http web framework for golang
- [goose](https://github.com/pressly/goose): db migration tool for golang

### Setup
Database setup
1. psql
2. create database botc
3. add migrations with goose: `goose -s create <file> sql
4. cd backend/db/migrations run `GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://<user>:<password>@localhost:5432/botc" goose up/down
