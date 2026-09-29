package services

import (
	"fins-api-services/db"
	"fins-api-services/structures"
	"log"
	"strconv"
)

func GetHistory(userId string) structures.Response {
	var response structures.Response

	records := db.GetHistory(userId)

	response.Data = records

	return response
}

func Expense(userId int, toUserId int, sum int) structures.Response {
	user := db.GetUserBalance(toUserId)
	var newSum int = user.Balance - sum
	var transactionResponse structures.Response

	log.Println(user.Balance)

	if user.Balance < sum {
		err := make(map[string]string)
		err["error"] = "На счету недостаточно средств, чтобы списать " + strconv.Itoa(sum)

		transactionResponse.ErrorText = err
		transactionResponse.Data = []structures.Transactions{}
	} else {
		record := db.TransactionProcess(userId, toUserId, newSum, 1, user.Balance)

		transactionResponse.Data = record
	}

	return transactionResponse
}

func Income(userId int, toUserId int, sum int) structures.Response {
	var response structures.Response

	user := db.GetUserBalance(toUserId)
	var newSum int = sum + user.Balance

	record := db.TransactionProcess(userId, toUserId, newSum, 2, user.Balance)

	response.Data = record

	return response
}
