package txtdb

import (
	"encoding/json"
	"errors"
    "fmt"
	"log"
	"time"

    "github.com/andersonlira/wallet-api/domain"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/str"
)

//GetExpenseList return all items 
func GetExpenseList() []domain.Expense {
	list := []domain.Expense{}
    fileName := fmt.Sprintf("bd/%ss.json", "Expense");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	return list
}

//GetExpenseByID return all items 
func GetExpenseByID(ID string) (domain.Expense, error) {
	list := GetExpenseList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Expense{}, errors.New("NOT_FOUND")
}



//SaveExpense saves a Expense object
func SaveExpense(it domain.Expense) domain.Expense {
	list := GetExpenseList()
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeExpense(list)
	return it
}

//UpdateExpense( updates a Expense object
func UpdateExpense(ID string, it domain.Expense) domain.Expense{
	list := GetExpenseList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writeExpense(list)
			return list[idx]
		}
	}
	return it
}

//DeleteExpense delete object by giving ID
func DeleteExpense(ID string) bool {
	list := GetExpenseList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeExpense(list)
			return true
		}
	}
	return false
}

func writeExpense(list []domain.Expense) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "Expense"), string(b))
}
