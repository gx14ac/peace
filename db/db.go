package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func MustNewDB(user, pass, name string) *gorm.DB {
	var db *gorm.DB
	var err error

	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		user,
		pass,
		name,
	)
	fmt.Println(url)

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
