package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/micro/v3/service/logger"
)

var (
	inited bool
	db     *sql.DB
	m      sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db has inited")
		logger.Log(1, err.Error())
		return
	}

	dsn := "micro:tF#262420228@tcp(127.0.0.1:3306)/user?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	inited = true
}

func GetDB() *sql.DB {
	if !inited {
		panic("DB does not init")
	}
	return db
}
