package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"time"

	"github.com/macfisherman/mboxreader/pkg/mboxreader"
)

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	mbox := mboxreader.NewMBOXReader(fh)
	for mbox.Next() {
		m, err := mail.ReadMessage(mbox)
		if err != nil {
			panic(err)
		}

		date, err := m.Header.Date()
		if err != nil {
			date = time.Unix(1594969314, 0)
		}
		fmt.Printf("%s\t%s\t%s\n", date.Format("2006-01-02 15:04 MST"), m.Header.Get("From"), m.Header.Get("Subject"))

		_, err = ioutil.ReadAll(m.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
