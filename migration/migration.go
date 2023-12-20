package migrations

import (
	"clinics/config"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

type MigrateCFG struct{}

var cfg = config.Load()

func init() {
	var (
		migrate     MigrateCFG
		databaseURL = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		)
	)

	fmt.Println("INIT FUNC")
	migrate.MigrateDatabaseUP(databaseURL)
	// migrate.MigrateDatabaseDOWN(databaseURL)
}
func (M *MigrateCFG) MigrateDatabaseUP(databaseURL string) {

	pathToMigrationDIR, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	pathToMigrationFiles := pathToMigrationDIR + "/database"
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFiles), databaseURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
		return
	}

	log.Println("migration done")

}

func (M *MigrateCFG) MigrateDatabaseDOWN(databaseURL string) {

	pathToMigrationDIR, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	pathToMigrationFiles := pathToMigrationDIR + "/database"
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFiles), databaseURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		log.Println(err)
		return
	}
	log.Println("migration DOWN...")
}
