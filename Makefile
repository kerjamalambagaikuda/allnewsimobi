dburl = mysql://root:secret@tcp(localhost:3356)/local
path =  $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# Database function
ifeq (db,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

db :
ifeq ($(RUN_ARGS), start)
	docker start mysql8
endif
ifeq ($(RUN_ARGS), createdb)
	docker exec -it mysql8 createdb --username=root --owner=root local
endif
ifeq ($(RUN_ARGS), dropdb)
	docker exec -it mysql8 dropdb local
endif
ifeq ($(RUN_ARGS), stop)
	docker stop mysql8
endif
ifeq ($(RUN_ARGS), migrate)
	migrate -path db/migration -database "${dburl}" -verbose up
endif
ifeq ($(RUN_ARGS), migrate-fresh)
	migrate -path db/migration -database "${dburl}" -verbose down
	migrate -path db/migration -database "${dburl}" -verbose up
endif

# Sql Function
ifeq (sqlc,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

sqlc:
ifeq ($(RUN_ARGS), init)
	docker run --rm -v ${path}:/src -w /src kjconroy/sqlc init
endif
ifeq ($(RUN_ARGS), generate)
	docker run --rm -v ${path}:/src -w /src kjconroy/sqlc generate
endif

# Sql Function
ifeq (server,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

server:
ifeq ($(RUN_ARGS), customer)
	go run ${path}cmd/customer/main.go
endif
ifeq ($(RUN_ARGS), status)
	go run ${path}cmd/customer/main.go
endif

# For Testing purpose
test:
	go test -v -cover ./...

# For Override same command
.PHONY: db sqlc test customer server