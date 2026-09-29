package db

import (
	"fins-api-services/structures"
	"log"
	"strconv"
)

func GetHistory(userId string) []structures.Transactions {
	var records []structures.Transactions

	rows, err := db.Query(`SELECT * FROM transactions WHERE user_id = ` + userId + ` ORDER BY date DESC`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var record structures.Transactions
		err = rows.Scan(&record.Id, &record.UserId, &record.AuthorId, &record.Type, &record.Sum, &record.Date)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records
}

func TransactionProcess(userId int, toUserId int, sum int, transactionType int, balance int) []structures.Transactions {
	var query string

	if transactionType == 1 {
		query = `
		INSERT INTO transactions (user_id, author_id, type, sum) 
		VAlUES (` + strconv.Itoa(userId) + `, ` + strconv.Itoa(toUserId) + `, ` + strconv.Itoa(transactionType) + `, ` + strconv.Itoa(balance-sum) + `)`

	} else {
		query = `
		INSERT INTO transactions (user_id, author_id, type, sum) 
		VAlUES (` + strconv.Itoa(userId) + `, ` + strconv.Itoa(toUserId) + `, ` + strconv.Itoa(transactionType) + `, ` + strconv.Itoa(sum-balance) + `)`
	}

	insertRow := db.QueryRow(query)

	updateRow := db.QueryRow(`UPDATE user SET balance = ` + strconv.Itoa(sum) + ` WHERE id = ` + strconv.Itoa(toUserId))

	log.Println(insertRow)
	log.Println(updateRow)

	return []structures.Transactions{}
}
