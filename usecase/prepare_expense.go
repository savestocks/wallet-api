package usecase

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
	"os"
	"errors"
	"strconv"
	"strings"
	"time"
	
	"github.com/andersonlira/wallet-api/domain"
	"github.com/andersonlira/wallet-api/gateway/txtdb"
)

func PrepareExpense(expense domain.Expense) (domain.Expense,error){
	account, err := findAccount(expense.Name)
	if err != nil {
		return expense,errors.New("not found")
	}
	expense.ID = account.ID
	expense.Name = account.Name
	txtdb.UpdateExpense(account.ID, expense)
	return expense,nil
}




func findAccount(accountNameOrID string) (account,error){
	gnc := loadXml()
	for i, account := range gnc.Book.Accounts {
		if gnc.Book.Accounts[i].Name == accountNameOrID  || gnc.Book.Accounts[i].ID == accountNameOrID{
			return  account,nil
		}
	}

	return account{}, errors.New("account not found")
}

func loadXml() (gnc gnc){
    // Open our xmlFile
    xmlFile, err := os.Open("bd/2020.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    // read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

    // we initialize our gnc array
    // we unmarshal our byteArray which contains our
    // xmlFiles content into 'gnc' which we defined above
    xml.Unmarshal(byteValue, &gnc)

    // we iterate through every user within our gnc array and
    // print out the user Type, their name, and their facebook url
	// as just an example

	return
}

func findransaction(gnc gnc,account account) int{
	tot := 0
	year, month, _ := time.Now().Date()
	base := fmt.Sprintf("%d-%02d",year,int(month))
	for _,t := range gnc.Book.Transactions {
		for _,s := range t.Splits.Splits {
			if s.Account == account.ID {
				idx := strings.Index(s.Value,"/")
				v,_ := strconv.Atoi(s.Value[0:idx])
				if strings.HasPrefix(t.DatePosted.Date,base) {
					tot+= v
				}

			}
		}
	}
	return tot
}


type gnc struct {
	Name xml.Name `xml:"gnc"`
	Book book `xml:"book"`
}
type book struct{
	XMLName xml.Name `xml:"book"`
	Accounts   []account   `xml:"account"`
	Transactions []transaction `xml:"transaction"`

}
type account struct{
	XMLName xml.Name `xml:"account"`
	ID string `xml:"id"`
	Name    string   `xml:"name"`
	

}
type transaction struct{
	XMLName xml.Name `xml:"transaction"`
	ID string `xml:"id"`
	DatePosted datePosted `xml:"date-posted"`
	Description string `xml:"description"`
	Splits splits `xml:"splits"`
}

type datePosted struct {
	XMLName xml.Name `xml:"date-posted"`
	Date string  `xml:"date"`

}



type splits struct{
	XMLName xml.Name `xml:"splits"`
	Splits []split `xml:"split"`
}
type split struct{
	XMLName xml.Name `xml:"split"`
	Account string `xml:"account"`
	Value string `xml:"value"`
}
