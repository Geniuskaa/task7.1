package transaction

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	//XMLName string `xml:"transaction"`
	Id      string  `json:"id"` //xml:"id"`
	From    string `json:"from"` //xml:"from"`
	To      string `json:"to"` //xml:"to"`
	Amount  int64  `json:"amount"` //xml:"amount"`
	Created int64  `json:"created"` //xml:"created"`
}

type Service struct {
	mu sync.Mutex
	transactions []*Transaction
}

func NewService() *Service {
	return &Service{}
}

type Writer struct {
	err error
	buf []byte
	n int
	wr io.Writer
}

func (s *Service) Register(from, to string, amount int64) (string, error) {

	t := &Transaction{
		Id:      "01",
		From:    from,
		To:      to,
		Amount:  amount,
		Created: time.Now().Unix(),
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions = append(s.transactions, t)

	return t.Id, nil
}

func (s *Service) Export(writer io.Writer) error {
	s.mu.Lock()
	if len(s.transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	records := make([][]string, len(s.transactions))
	for _, t := range s.transactions {
		record := []string {
			t.Id,
			t.From,
			t.To,
			strconv.FormatInt(t.Amount, 10),
			strconv.FormatInt(t.Created, 10),
		}
		records = append(records, record)
	}
	s.mu.Unlock()

	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func (s *Service) Import(filename string) ([]Transaction,error) {
	// data - []byte
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil,err
	}

	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil,err
	}

	sliceOfTransactions := make([]Transaction, 0)

	for _, m := range records {
		fmt.Println(m)  //- это для примера, непонятно почему выводит по другому
		sliceOfTransactions = append(sliceOfTransactions, mapRowToTransaction(m))
	}
	return sliceOfTransactions, nil
}

//////Тут json

func ExportJson(sliceOfTransactions []Transaction)  ([]byte, error) { //sliceOFTransactions []Transaction

	encoded, err := json.Marshal(sliceOfTransactions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(encoded))

	/*var decoded []Transaction
	// Важно: передаём указатель, чтобы функция смогла записать данные
	err = json.Unmarshal(encoded, &decoded)
	log.Printf("%#v", decoded)*/


	return encoded, nil
}


func ImportJson(encoded []byte) ([]Transaction, error) {
	var decoded []Transaction
	// важно: передаём указатель
	err := json.Unmarshal(encoded, &decoded)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("%#v", decoded)

	return decoded,nil
}




func mapRowToTransaction(slice []string) Transaction{
	var a Transaction
	for i := 0; i <= 4; i++ {
		switch i {
		case 0:
			a.Id = slice[i]
			break
		case 1:
			a.From = slice[i]
			break
		case 2:
			a.To = slice[i]
			break
		case 3:
			x, _ := strconv.Atoi(slice[i])
			a.Amount = int64(x)
			break
		case 4:
			x, _ := strconv.Atoi(slice[i])
			a.Created = int64(x)
			break
		default:
			break
		}
	}
	return a
}
