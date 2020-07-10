// Copyright (c) 2020 Jeff Macdonald <macfisherman@gmail.com>
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

package mboxreader

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/textproto"
	"regexp"
)

type MBox struct {
	t *textproto.Reader
	//	buffer bytes.Buffer
	//line []byte
	err error
}

var escapedFrom = regexp.MustCompile(`>(>*From )`)

func NewMBOXReader(r io.Reader) (*MBox, error) {
	t := textproto.NewReader(bufio.NewReader(r))
	line, err := t.ReadLineBytes()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(line[:5], []byte("From ")) {
		return nil, errors.New("Missing mbox header")
	}

	return &MBox{
		t: t,
		//		line: line,
	}, nil
}

func (m *MBox) Read(p []byte) (int, error) {
	line, err := m.t.ReadLineBytes()
	if err != nil {
		m.err = err
		// copy what there might be to the caller
		if len(line) > 0 {
			n := copy(p, line)
			return n, err
		}

		return 0, err
	}

	// blank line, is it the message terminator?
	if len(line) == 0 {
		demark, err := m.t.R.Peek(5)
		if (err == io.EOF) && (len(demark) == 0) {
			m.err = err
			return 0, err
		}
		if (err == io.EOF) && (len(demark) > 0) {
			m.err = err
			if demark[len(demark)-1] != '\n' {
				return 0, errors.New("mbox file does not end in a newline")
			}

			n := copy(p, demark[:len(demark)])
			return n, err
		}
		if err != nil {
			m.err = err
			return 0, io.EOF
		}

		// if the blank line is followed by
		// a message header, then the end
		// of the message has been reached
		if bytes.Equal(demark, []byte("From ")) {
			_, err := m.t.ReadLineBytes() // throw the header away
			m.err = err
			return 0, io.EOF
		}
	}

	// remove leading > from an escaped From
	if escapedFrom.Match(line) {
		line = line[1:]
	}

	line = append(line, '\n')
	n := copy(p, line)

	return n, nil
}

func (m *MBox) Next() bool {
	return m.err != io.EOF
}
