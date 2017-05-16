// mysql.go
// author:昌维 [github.com/cw1997]
// date:2017-05-08 01:36:41
package db

import (
	"database/sql"
	//	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"config"
)

var (
	db *sql.DB
)

type DatabaseConfig struct {
	ip          string
	port        string
	username    string
	password    string
	database    string
	charset     string
	prefix      string
	maxIdle     int
	maxOpen     int
	maxLifetime time.Duration
}

func ConnectPool() (*sql.DB, error) {
	//	db, err := sql.Open("mysql", "root:password@(127.0.0.1)/dbname?charset=utf8")
	maxIdle, _ := strconv.Atoi(config.Get("database.maxIdle"))
	maxOpen, _ := strconv.Atoi(config.Get("database.maxOpen"))
	maxLifetime, _ := strconv.Atoi(config.Get("database.maxLifetime"))
	dbconfig := DatabaseConfig{
		ip:          config.Get("database.ip"),
		port:        config.Get("database.port"),
		username:    config.Get("database.username"),
		password:    config.Get("database.password"),
		database:    config.Get("database.database"),
		charset:     config.Get("database.charset"),
		prefix:      config.Get("database.prefix"),
		maxIdle:     maxIdle,
		maxOpen:     maxOpen,
		maxLifetime: time.Duration(maxLifetime) * time.Second,
	}
	uri := dbconfig.username + ":" + dbconfig.password + "@tcp(" + dbconfig.ip + ":" + dbconfig.port + ")/" + dbconfig.database + "?charset=" + dbconfig.charset + "&allowOldPasswords=1"
	//	fmt.Println(uri)
	dbconnect, err := sql.Open("mysql", uri)
	dbconnect.SetMaxIdleConns(dbconfig.maxIdle)
	dbconnect.SetMaxOpenConns(dbconfig.maxOpen)
	dbconnect.SetConnMaxLifetime(dbconfig.maxLifetime)
	db = dbconnect
	return db, err
}

func connect() *sql.DB {
	return db
}

//func GetOne(sql string, param ...string) {
//	db, err := Connect()
//	stmt, err := db.Prepare(sql)
//	if stmt != nil {
//		rows, err := stmt.Query(param...)
//		if err != nil {
//			fmt.Println(rows)
//		}
//		stmt.Close()
//	}
//}
func GetLongUrlByShortUrl(shortUrl string) (string, error) {
	var longUrl string
	row := db.QueryRow("SELECT `longurl` FROM `"+config.Get("database.prefix")+"url` WHERE `shorturl` = ?", shortUrl)
	//	defer row.Close()
	rowErr := row.Scan(&longUrl)
	if rowErr != nil {
		//		fmt.Printf("get row error: %v\n", rowErr)
		return "", rowErr
	}
	//	fmt.Println(longUrl)
	return longUrl, nil
}

//func GetShortUrlByLongUrl(longUrl string) (string, error) {
//	var shortUrl string
//	row := db.QueryRow("SELECT `shorturl` FROM `"+config.Get("database.prefix")+"url` WHERE `longurl` = ?", longUrl)
//	//	defer row.Close()
//	rowErr := row.Scan(&shortUrl)
//	if rowErr != nil {
//		log.Printf("get row error: %v\n", rowErr)
//		return "", rowErr
//	}
//	//	fmt.Println(longUrl)
//	return shortUrl, nil
//}

func SetShortUrlByLongUrl(longUrl string, shortUrl string, datetime string, ip string) bool {
	db := connect()
	ret, err := db.Exec("INSERT INTO `tinyurl_url` (`longurl`, `shorturl`, `add_time`, `add_ip`) VALUES (?, ?, ?, ?)", longUrl, shortUrl, datetime, ip)
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return false
	}
	//	if LastInsertId, err := ret.LastInsertId(); nil == err {
	//		fmt.Println("LastInsertId:", LastInsertId)
	//	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		//		fmt.Println("RowsAffected:", RowsAffected)
		return RowsAffected > 0
	}
	return false
}

func Execute(sql string) error {
	db := connect()
	_, err := db.Exec(sql)
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return err
	}
	return nil
}
