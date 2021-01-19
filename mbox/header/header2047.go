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

package header

import (
	"fmt"
	"io"
	"mime"
	"net/mail"
	"net/textproto"
	"time"

	"github.com/macfisherman/mboxreader/mbox/header/charmap"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
)

var alias = map[string]string{
	"gb2312":         "gbk",
	"cp1252":         "windows-1252",
	"cp1256":         "windows-1256",
	"x-gbk":          "gbk",
	"unknown-8bit":   "windows-1252",
	"ks_c_5601-1987": "euc-kr",
	"shift-jis":      "shift_jis",
}

var localCharmap = map[string]encoding.Encoding{
	"roman8":    charmap.HPRoman8,
	"hp-roman8": charmap.HPRoman8,
}

type Header struct {
	header  mail.Header
	decoder *mime.WordDecoder
}

func New(h mail.Header) Header {
	dec := new(mime.WordDecoder)
	dec.CharsetReader = charsetReader
	return Header{header: h, decoder: dec}
}

func (h Header) AddressList(key string) ([]*mail.Address, error) {
	return h.header.AddressList(key)
}

func (h Header) Date() (time.Time, error) {
	return h.header.Date()
}

func (h Header) Get(key string) (string, error) {
	raw := h.header.Get(key)
	return h.decode(raw)
}

func (h Header) Values(key string) ([]string, error) {
	var values []string

	rawValues := textproto.MIMEHeader(h.header).Values(key)
	for _, rawValue := range rawValues {
		value, err := h.decode(rawValue)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func (h Header) decode(raw string) (string, error) {
	header, err := h.decoder.DecodeHeader(raw)
	if err != nil {
		return "", err
	}

	return header, nil
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	e, err := findEncoder(charset)
	if err != nil {
		return nil, err
	}

	return e.NewDecoder().Reader(input), nil
}

func withAlias(charset string) string {
	if a, ok := alias[charset]; ok {
		return a
	}

	return charset
}

func haveLocal(charset string) encoding.Encoding {
	if l, ok := localCharmap[charset]; ok {
		return l
	}

	return nil
}

func findEncoder(charset string) (encoding.Encoding, error) {
	if e := haveLocal(charset); e != nil {
		return e, nil
	}

	e, err := ianaindex.MIME.Encoding(withAlias(charset))
	if (e == nil) && (err == nil) {
		return nil, fmt.Errorf("%s not supported", charset)
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, charset)
	}

	return e, nil
}
