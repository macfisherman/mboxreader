// Copyright (c) 2021 Jeff Macdonald <macfisherman@gmail.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/macfisherman/mboxreader/mbox"
	"github.com/macfisherman/mboxreader/mbox/header"
)

func main() {
	log.SetFlags(log.Lshortfile)

	fh, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	index := 0
	box := mbox.New(fh)
	for box.Next() {
		m, err := mail.ReadMessage(box)
		if err != nil {
			log.Printf("skipping message at %d: got premature %v", index, err)
			index++
			continue
		}
		h := header.New(m.Header)

		date, err := h.Date()
		if err != nil {
			date = time.Unix(1594969314, 0)
		}
		fmt.Print(date.Format("2006-01-02 15:04 MST"))

		from, err := h.Get("From")
		if err != nil {
			log.Fatal("From: ", err)
		}
		fmt.Print("\t", from)

		subject, err := h.Get("Subject")
		subject = strings.ReplaceAll(subject, "\n", "↵")
		if err != nil {
			log.Fatal("\n", err)
		}
		fmt.Print("\t", subject, "\n")

		_, err = ioutil.ReadAll(m.Body)
		if err != nil {
			log.Fatal("body: ", err)
		}

		index++
	}
}
