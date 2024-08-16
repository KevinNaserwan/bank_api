package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinnaserwan/bank_api/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()
	testDB, err = pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
