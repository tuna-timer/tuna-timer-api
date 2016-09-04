package utils

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/olebedev/config"

	"fmt"
	// a DB package
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"github.com/tanel/dbmigrate"
)

const (
	ProductionEnv		= "production"
	TestEnv			= "test"
)

const (
	ConfigFile string 	= "../config.yml"
	MigrationsFolder string = "../data/migrations/"
)

// Environment is a thing that holds env. specific stuff
type Environment struct {
	OrmDB *gorm.DB
	RawDB *sql.DB // for unit tests
}

// NewEnvironment creates a new environment
func NewEnvironment(environment string) (*Environment, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}
	cfg.Env() // test this out!
	cfg, err = cfg.Get(environment)

	env := &Environment{}
	connection, err := connectToDatabase(cfg)
	if err != nil {
		return nil, err
	}
	env.OrmDB = connection
	env.RawDB = env.OrmDB.DB()
	return env, nil
}

func (env *Environment) ReleaseResources() {
	log.Println("*********************** ReleaseResources")
	env.OrmDB.Close()
}

func (env *Environment) MigrateDatabase() error {
	log.Println("Migrating database...")

	err := dbmigrate.Run(env.RawDB, MigrationsFolder)
	if err != nil {
		return err
	}

	log.Println("Database migrated!")
	return nil
}

func connectToDatabase(cfg *config.Config) (*gorm.DB, error) {

	log.Println("Connecting to database:")
	log.Printf("database.name: %s", cfg.UString("database.name"))
	log.Printf("database.host: %s", cfg.UString("database.host"))
	log.Printf("database.port: %s", cfg.UString("database.port"))
	log.Printf("database.user: %s", cfg.UString("database.user"))
	log.Print("database.pass: ***********")

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("sslmode=disable dbname=%s host=%s port=%s user=%s password=%s",
			cfg.UString("database.name"),
			cfg.UString("database.host"),
			cfg.UString("database.port"),
			cfg.UString("database.user"),
			cfg.UString("database.pass"),
		))

	if err != nil {
		log.Println("Failed to connect!!")
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected successfully!")
	return db, nil
}

func readConfig() (*config.Config, error) {
	cfg, err := config.ParseYamlFile(ConfigFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
