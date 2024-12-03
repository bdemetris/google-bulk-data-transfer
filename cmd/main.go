package main

import (
	"fmt"
	"log"
	"os"

	transfer "github.com/bdemetris/google-bulk-drive-transfer/google"
	"github.com/gocarina/gocsv"
)

const (
	csvFile         = "users.csv"
	credentialsFile = "credentials.json"
	domainDelegate  = "delegate@domain.com"
)

type UserPair struct {
	Source      string `csv:"source_email"`
	Destination string `csv:"destination_email"`
}

func main() {
	client, err := transfer.NewService(credentialsFile, domainDelegate)
	if err != nil {
		fmt.Println(err)
	}

	users, err := readCsvFile(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	for _, u := range users {
		if err := client.RequestTransfer(u.Source, u.Destination); err != nil {
			fmt.Println(err)
			continue
		}
	}

}

func readCsvFile(filePath string) ([]*UserPair, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	users := []*UserPair{}

	if err := gocsv.UnmarshalFile(f, &users); err != nil {
		return nil, err
	}

	return users, nil
}
