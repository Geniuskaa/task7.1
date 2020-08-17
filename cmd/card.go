package main

import (
	"fmt"
	"io"
	"log"
	"os"
	transaction "task7.1/pkg/transactions"
)

func main() {
	/*if err := execute("export.csv"); err != nil {
		os.Exit(1)
	}*/
	svc := transaction.NewService()
	fmt.Println(svc.Import("export.csv"))
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
