package util

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sQLConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Dbname    string `json:"dbname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

func MigrationDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: MigrationPath,
	}

	n, err := migrate.Exec(sqlDB, DriverName, migrations, migrate.Up)

	if err != nil {
		return err
	}

	n, err = migrate.Exec(sqlDB, DriverName, migrations, migrate.Up)

	if err != nil {
		return err
	}

	fmt.Println("Applied migraions! ", n)

	return nil
}

func loadDBConfigs(filepath string) *sQLConfig {
	configFile, err := os.Open(filepath)

	defer configFile.Close()

	if err != nil {
		return nil
	}

	jsonParser := json.NewDecoder(configFile)
	config := sQLConfig{}
	jsonParser.Decode(&config)

	return &config
}

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")
	parseTime := os.Getenv("DB_PARSE_TIME")
	loc := os.Getenv("DB_LOCAL")

	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", userName, password, host, port, dbName, charset, parseTime, loc)

	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	//Turn on debug log
	db = db.Debug()

	//Migration & Seed into database
	if err := MigrationDB(db); err != nil {
		fmt.Println(fmt.Sprint(err))
		return nil
	}

	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
