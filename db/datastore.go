package db

import (
	"ginTest/config"
	"fmt"
	"net"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func InitDB()  {

	conf := config.Config{}

	addr := conf.DBConfig.DBAddr
	port := conf.DBConfig.DBPort
	username := conf.DBConfig.DBUsername
	password := conf.DBConfig.DBPasseord
	base := conf.DBConfig.DBBase

	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",username, password, addr, port, base)

	ch := make(chan int, 1)

	go func() {
		var err error
		var c net.Conn
		c, err = net.DialTimeout("tcp", addr+":"+port, 20*time.Second)
		if err == nil {
			c.Close()
			ch <- 1
		} else {
			log.Errorf("Failed to connect to db, retry after 2 seconds: v%", err)
			time.Sleep(2 * time.Second)
		}
	}()

	select {
	case <- ch:
	case <-time.After(60 * time.Second):
		panic("Failed to connect to DB after 60s")
	}

	var err error
	db, err = gorm.Open("mysql", dbStr)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(1)
	db.DB().SetMaxOpenConns(2)
	db.LogMode(true)

}
