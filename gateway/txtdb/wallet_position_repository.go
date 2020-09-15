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
	"sort"
)

//GetWalletPositionList return all items 
func GetWalletPositionList() []domain.WalletPosition {
	list := []domain.WalletPosition{}
    fileName := fmt.Sprintf("bd/%ss.json", "WalletPosition");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)

	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list
}

//GetWalletPositionByExpenseID return all items 
func GetLastWalletPositionByExpenseID(ID string) (domain.WalletPosition, error) {
	list := GetWalletPositionList()
	for idx, _ := range list {
		if(list[idx].ExpenseID == ID){
			return list[idx],nil
		}
	}
	return domain.WalletPosition{}, errors.New("NOT_FOUND")
}



//SaveWalletPosition saves a WalletPosition object
func SaveWalletPosition(it domain.WalletPosition) domain.WalletPosition {
	list := GetWalletPositionList()
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeWalletPosition(list)
	return it
}

//UpdateWalletPosition( updates a WalletPosition object
func UpdateWalletPosition(ID string, it domain.WalletPosition) domain.WalletPosition{
	list := GetWalletPositionList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writeWalletPosition(list)
			return list[idx]
		}
	}
	return it
}

//DeleteWalletPosition delete object by giving ID
func DeleteWalletPosition(ID string) bool {
	list := GetWalletPositionList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeWalletPosition(list)
			return true
		}
	}
	return false
}

func writeWalletPosition(list []domain.WalletPosition) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "WalletPosition"), string(b))
}



