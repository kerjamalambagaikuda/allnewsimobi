# ALL NEW SIMOBI

# Installation

**General :**

1. Install [Golang](https://golang.org/dl/)

2. Install [Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
   For [Windows OS](#windows) install this additional first.

3. Install [SqlC](https://docs.sqlc.dev/en/latest/overview/install.html) or [for Windows](https://github.com/kyleconroy/sqlc/releases/download/v1.10.0/sqlc_1.10.0_windows_amd64.zip).
   If blokced by TM APEX Antivirus try install via [Docker](#sqlc)

<!----><a id="windows"></a>

**Additional for Windows:**

- Install [Scoop](https://scoop.sh/)

- Install [Cygwin](http://www.cygwin.com/)

---

# Migration

Create migration file, it will create .down.sql and .up.sql file in folder db/migration

```
migrate create -ext sql -dir db/migration -seq name_table
```

## Run migration

Migrate without drop table

```
make migrate
```

Migrate with drop table

```
make migrate-fresh
```

---

# Configuration

<!----><a id="sqlc"></a>

**SQLC**

Install sqlc via docker, then update your path project in Makefile

```
docker pull kjconroy/sqlc
```

---

**Database**

- Update your database complete url in Makefile
- Update your database ur, user, password in app.env
- Install image mysql in docker, or follow [This Instruction](https://hub.docker.com/_/mysql)

# How to run

Run docker app, then type

```
make docker start
```
