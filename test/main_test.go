package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	sqlc "github.com/kerjamalambagaikuda/allnewsimobi/db/sqlc"
	utilConfig "github.com/kerjamalambagaikuda/allnewsimobi/util/config"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load Config
	config, err := utilConfig.LoadConfig("../")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Load Logger
	err = utilConfig.LoadLogger("testing", config)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Test DB Connection
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
