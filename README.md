# ALL NEW SIMOBI
New Service for Simobi
  - [Installation](#installation)
  - [Migration](#migration)
  - [Configuration](#configuration)
  - [Command](#command)
  - [Todos](#todo)
---
## Installation
### **General**
1. Install [Golang](https://golang.org/dl/)
2. Install [Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
   For [Windows OS](#additional-for-windows) install this additional first.
3. Install [SqlC](https://docs.sqlc.dev/en/latest/overview/install.html) or [for Windows](https://github.com/kyleconroy/sqlc/releases/download/v1.10.0/sqlc_1.10.0_windows_amd64.zip).
   If blokced by TM APEX Antivirus try install via [Docker](#sqlc)
### **Additional for Windows**
- Install [Scoop](https://scoop.sh/)
- Install [Cygwin](http://www.cygwin.com/)
---
## Migration
Create migration file, it will create .down.sql and .up.sql file in folder db/migration
```
migrate create -ext sql -dir db/migration -seq name_table
```
---
## Configuration
### **SQLC**
Install sqlc via docker, then update your path project in Makefile
```
docker pull kjconroy/sqlc
```
### **Database**
- Update your database complete url in Makefile
- Update your database url, user, password in app.env
- Install image mysql in docker, or follow [This Instruction](https://hub.docker.com/_/mysql)
---
## Commands
Initialize module go
```
go mod init name_module
```
Get library for go
```
go get library_github_url
```
Clean unused library go
```
go mod tidy
```
Run mysql container in docker
```
make db start
```
Migrate without drop table
```
make migrate
```
Migrate with drop table
```
make migrate-fresh
```
Generate query sqlc
```
make sqlc generate
```
Run service
```
make server customer
```
Testing
```
make test
```
---
## Todos
- [x] Migration
- [x] SQLC
- [ ] Trx
- [x] Logger
- [ ] Testing
- [ ] Makefile
- [ ] Dockerfile
- [ ] gRPC
- [ ] Rest API
- [ ] Authorization
