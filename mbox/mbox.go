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

package mbox

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"regexp"
)

type state int

const (
	HEADER state = iota
	FOOTER
)

type MBox struct {
	r   *bufio.Reader
	buf bytes.Buffer
	err error
}

var escapedFrom = regexp.MustCompile(`>(>*From )`)

func New(r io.Reader) *MBox {
	return &MBox{
		r: bufio.NewReader(r),
	}
}

func (m *MBox) Read(p []byte) (int, error) {
	n, err := m.buf.Read(p)
	if err != nil {
		return n, err
	}

	return n, nil
}

func (m *MBox) Next() bool {
	if m.err == io.EOF {
		return false
	}

	if err := consumeHeader(m.r); err != nil {
		m.err = err
		return false
	}

	if err := m.getMessage(); err != nil {
		m.err = err
		if err == io.EOF {
			return true
		}
		return false
	}

	return m.err == nil
}

func (m *MBox) getMessage() error {
	for {
		line, err := m.r.ReadBytes('\n')
		if err != nil {
			return err
		}

		if bytes.Equal(line, []byte("\n")) ||
			bytes.Equal(line, []byte("\r\n")) {
			nextLine, err := m.r.Peek(5) // Looking for From<space>
			if (err != nil) && (err != io.EOF) {
				return err
			}
			if bytes.HasPrefix(nextLine, []byte("From ")) {
				return nil
			}
		}

		if escapedFrom.Match(line) {
			line = line[1:]
		}

		_, err = m.buf.Write(line)
		if err != nil {
			return err
		}

	}
}

func consumeHeader(r *bufio.Reader) error {
	line, err := r.ReadBytes('\n')
	if err != nil {
		return err
	}

	if !bytes.HasPrefix(line, []byte("From ")) {
		return errors.New("header does not start with 'From '")
	}

	return nil
}

func (m MBox) Error() error {
	return m.err
}
