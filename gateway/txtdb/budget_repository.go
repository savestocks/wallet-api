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

//GetBudgetList return all items 
func GetBudgetList() []domain.Budget {
	list := []domain.Budget{}
    fileName := fmt.Sprintf("bd/%ss.json", "Budget");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	return list
}

//GetBudgetByID return all items 
func GetBudgetByID(ID string) (domain.Budget, error) {
	list := GetBudgetList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Budget{}, errors.New("NOT_FOUND")
}



//SaveBudget saves a Budget object
func SaveBudget(it domain.Budget) domain.Budget {
	list := GetBudgetList()
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeBudget(list)
	return it
}

//UpdateBudget( updates a Budget object
func UpdateBudget(ID string, it domain.Budget) domain.Budget{
	list := GetBudgetList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writeBudget(list)
			return list[idx]
		}
	}
	return it
}

//DeleteBudget delete object by giving ID
func DeleteBudget(ID string) bool {
	list := GetBudgetList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeBudget(list)
			return true
		}
	}
	return false
}

func writeBudget(list []domain.Budget) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "Budget"), string(b))
}
