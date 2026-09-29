package db

import (
	"fins-api-services/structures"
	"log"
	"strconv"
)

func GetUserBalance(userId int) structures.UserBalance {
	var userBalance structures.UserBalance

	log.Println(userId)

	log.Println("SELECT id, balance FROM user WHERE id = " + strconv.Itoa(userId))

	row := db.QueryRow(`SELECT id, balance FROM user WHERE id = ` + strconv.Itoa(userId))

	log.Println(row)

	errors := row.Scan(&userBalance.Id, &userBalance.Balance)

	if errors != nil {
		log.Println(errors)
	}

	return userBalance
}
