package main

import (
	"io"
	"log"
	"os"
	transaction "task7.1/pkg/transactions"
	"time"
)

func main() {
	/*if err := execute("export3.json"); err != nil {
		os.Exit(1)
	}*/
	//svc := transaction.NewService()
	//fmt.Println(svc.Import("export.csv"))

	transactions := []transaction.Transaction{
		{
			Id:      "1",
			From:    "0001",
			To:      "0002",
			Amount:  100_00,
			Created: time.Now().Unix(),
		},
		{
			Id:      "2",
			From:    "0001",
			To:      "0002",
			Amount:  200_00,
			Created: time.Now().Unix(),
		},
	}

	file, _ := transaction.ExportJson(transactions)
	slice, _ := transaction.ImportJson(file)
	log.Println(slice)

}

	//transaction.ExportJson(transactions)
	//fileJsonXml("export3.csv", transactions)




func fileJsonXml (filename string, sliceOfTransactions []transaction.Transaction) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	//svc := transaction.NewService()

	//err = transaction.ExportJson(sliceOfTransactions)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}

func execute(filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	svc := transaction.NewService()

	_, err = svc.Register("0001","0002",10_000_00)
	if err != nil {
		log.Println(err)
		return
	}

	err = svc.Export(file)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}
