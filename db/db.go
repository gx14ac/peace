package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func MustNewDB(host, port, user, pass, name string) *gorm.DB {
	var db *gorm.DB
	var err error

	url := fmt.Sprintf("%s:%s@tcp([%s]:%s)/%s?charset=utf8&parseTime=true",
		user,
		pass,
		host,
		port,
		name,
	)

	x := 1
	maxRetryCount := 15

	for {
		fmt.Println("Trying to connect DB...", url)

		if db, err = gorm.Open("mysql", url); err != nil {
			fmt.Println("Something wrong with connecting DB...", err.Error())
		} else {
			fmt.Println("Succeeded!")
			break
		}

		time.Sleep(1000 * time.Millisecond)

		if x < maxRetryCount {
			x++
			continue
		} else {
			panic("Failed to connect DB")
		}
	}

	return db
}
