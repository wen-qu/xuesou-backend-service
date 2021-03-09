package basic

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/util/log"
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
		log.Logf(err.Error())
		return
	}

	dsn := "micro:tF#262420228@tcp(127.0.0.1:3306)/test?charset=utf8"
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
