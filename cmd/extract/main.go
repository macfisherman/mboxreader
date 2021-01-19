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
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/macfisherman/mboxreader/mbox"
)

func main() {
	term := flag.String("term", "", "search term")
	count := flag.Int("count", 1, "N messages")
	offset := flag.Int("msg", -1, "get N msg")
	scan := flag.Bool("scan", false, "scan mbx, report msg count")

	flag.Parse()

	fh, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	index := 0
	mbox := mbox.New(fh)
	for mbox.Next() {
		msg, err := ioutil.ReadAll(mbox)
		if err != nil {
			log.Fatal(err)
		}

		if *term != "" {
			if bytes.Contains(msg, []byte(*term)) {
				fmt.Println(string(msg))
				*count--
				if *count == 0 {
					break
				}
			}
		}

		if index == *offset {
			fmt.Println(string(msg))
			break
		}

		index++
	}

	if *scan {
		fmt.Println("Msgs: ", index)
	}

	if mbox.Error() != nil {
		log.Fatal(mbox.Error())
	}
}
