package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

var (
	dbURIs []string = []string{
		"postgresql://postgres:12345@db1:5432/load_balancer?sslmode=disable",
		"postgresql://postgres:12345@db2:5432/load_balancer?sslmode=disable",
	}
	dbConns []*sqlx.DB = []*sqlx.DB{}
	i       int        = 0
)

const (
	dbDriver      = "postgres"
	migrationPath = "./app/migrations"
)

func GetDB() *sqlx.DB {
	if i == len(dbConns) {
		i = 0
	}

	db := dbConns[i]
	i += 1
	return db

}

func GetDBByID(id int) *sqlx.DB {
	if id%2 == 0 {
		return dbConns[1]
	}
	return dbConns[0]
}

func SetDB() {
	for i, uri := range dbURIs {
		db, err := InitDB(uri, i+1)
		if err != nil {
			panic("error while establishing database connections")
		}
		dbConns = append(dbConns, db)
	}
}

func InitDB(uri string, i int) (db *sqlx.DB, err error) {
	conn, err := sqlx.Connect(dbDriver, uri)
	if err != nil {
		log.Println("InitDB: ", err.Error())
		return
	}

	err = RunMigrations(uri, i)
	if err != nil {
		log.Println("InitDB: ", err.Error())
		return
	}
	return conn, nil
}

func RunMigrations(uri string, i int) (err error) {
	db, err := sql.Open(dbDriver, uri)
	if err != nil {
		return
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println("RunMigrations: ", err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance(getMigrationPath()+"/"+strconv.Itoa(i), dbDriver, driver)
	if err != nil {
		log.Println("RunMigrations: ", err.Error())
		return
	}

	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		err = nil
		return
	}
	return
}

func getMigrationPath() string {
	return fmt.Sprintf("file://%s", migrationPath)
}
