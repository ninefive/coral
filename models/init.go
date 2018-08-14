package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self *gorm.DB
	//	Docker *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s", username, password, addr, name, true, "Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//设置最大连接数，默认值为0表示不限制。
	//设置最大连接数可以避免并发太高导致连接mysql出现too many connection的错误
	db.DB().SetMaxOpenConns(5000)
	//设置闲置连接数。
	//当开启的一个连接使用完成后可以放在连接池里等候下次使用。
	db.DB().SetMaxIdleConns(50)
}

func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

/*
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"),
	)
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}
*/

func (db *Database) Init() {
	DB = &Database{
		Self: InitSelfDB(),
		//		Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	//	DB.Docker.Close()
}
