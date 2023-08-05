package main

import (
	"AuthWithJWT"
	"AuthWithJWT/migrations"
	"AuthWithJWT/pkg/handler"
	"AuthWithJWT/pkg/repository"
	"AuthWithJWT/pkg/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loadong env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	runMigrate("postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(AuthWithJWT.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func runMigrate(dsn string) {
	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(dsn, s)
}

func runDBMigrate(dsn string, source *bindata.AssetSource) {
	d, err := bindata.WithInstance(source)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
