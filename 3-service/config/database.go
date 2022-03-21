package config

import (
	"fmt"
	"navigate/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 链接数据库
func SetUpDatabaseConnection() *gorm.DB {

	// dbUser := os.Getenv("DB_USER")
	// dbHost := os.Getenv("DB_HOST")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbPort := os.Getenv("DB_PORT")
	// println(dbUser)

	// // DSN:               "root:123456@tcp(127.0.0.1:3306)/yzy_navigate?charset=utf8mb4&parseTime=True&loc=Local",
	// dsnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/yzy_navigate?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "yzy_", // 表前缀
		},
	})

	if err != nil {
		fmt.Sprintln("database connect error")
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.ResourceType{})
	return db
}

// 关闭数据库
func TearDownDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		fmt.Println("database close to form ")
	}
	dbSql.Close()
}
