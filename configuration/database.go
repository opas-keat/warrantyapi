package configuration

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"warrantyapi/entity"
	"warrantyapi/exception"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, _ := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, _ := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	dsn := "user=" + username + " password=" + password + " dbname=" + dbName + " host=" + host + " port=" + port + " TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		// CreateBatchSize: 100,
		// NamingStrategy: schema.NamingStrategy{
		// 	// TablePrefix:   "ect_",                            // table name prefix, table for `User` would be `t_users`
		// 	SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
		// 	NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		// },
		Logger: newLogger,
	})
	// exception.PanicLogging(err)
	// db.LogMode(true)

	sqlDB, _ := db.DB()
	exception.PanicLogging(err)
	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	err = db.AutoMigrate(&entity.Dealer{})
	if err != nil {
		fmt.Println("failed to automigrate Dealer entity:", err.Error())
		// return db
	}
	err = db.AutoMigrate(&entity.Log{})
	if err != nil {
		fmt.Println("failed to automigrate Log entity:", err.Error())
		// return db
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		fmt.Println("failed to automigrate Product entity:", err.Error())
		// return db
	}
	err = db.AutoMigrate(&entity.Warranty{})
	if err != nil {
		fmt.Println("failed to automigrate Warranty entity:", err.Error())
		// return db
	}
	// err = db.AutoMigrate(&entity.Customer{})
	// if err != nil {
	// 	fmt.Println("failed to automigrate Customer entity:", err.Error())
	// 	return db
	// }
	// err = db.AutoMigrate(&entity.Warranty{})
	// if err != nil {
	// 	fmt.Println("failed to automigrate Warranty entity:", err.Error())
	// 	return db
	// }
	return db
}
