package usecase

import (
	"time"
	"fmt"
	"github.com/andersonlira/wallet-api/domain"
	"github.com/andersonlira/wallet-api/gateway/txtdb"
)

func GetWalletPositionByExpenseID(ID string) (domain.WalletPosition, error) {
	wallet, err := txtdb.GetLastWalletPositionByExpenseID(ID)
	past := time.Now().Add(-3*time.Hour)
	fmt.Println(past)
	if err != nil || wallet.CreatedAt.Before(past) {
		account, err := findAccount(ID)
		if err != nil {
			return wallet, err
		}
		wallet = txtdb.SaveWalletPosition(prepareWallet(account))
	}
	return wallet, nil
}

func prepareWallet(account account) domain.WalletPosition {
	wallet := domain.WalletPosition{}
	gnc := loadXml()
	wallet.CreatedAt = time.Now()
	wallet.Total = findransaction(gnc,account)
	wallet.ExpenseID = account.ID
	return wallet
}