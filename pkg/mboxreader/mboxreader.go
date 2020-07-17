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
	"regexp"
)

type State int

const (
	HEADER State = iota
	EMPTYLINE
	FOOTERHEADER
)

type MBox struct {
	r         *bufio.Reader
	err       error
	buffer    []byte
	offset    int
	initState State
}

var escapedFrom = regexp.MustCompile(`>(>*From )`)

func NewMBOXReader(r io.Reader) *MBox {
	return &MBox{
		r:         bufio.NewReader(r),
		initState: HEADER,
	}
}

func (m *MBox) Read(p []byte) (int, error) {
	if (m.buffer != nil) && (m.offset >= len(m.buffer)) {
		m.buffer = nil
		m.offset = 0
		return 0, io.EOF
	}

	var err error
	if m.buffer == nil {
		var msg []byte
		msg, err = readMessage(m.r, m.initState)
		if err == io.EOF {
			m.err = io.EOF
		}

		if (err != nil) && (err != io.EOF) {
			return 0, err
		}

		m.buffer = msg
		m.initState = EMPTYLINE
	}

	n := copy(p, m.buffer[m.offset:])
	m.offset += n

	return n, err
}

func readMessage(r *bufio.Reader, state State) ([]byte, error) {
	var msg []byte
	var eol []byte

	for {
		line, err := r.ReadBytes('\n')
		if (err != nil) && (err != io.EOF) {
			return nil, err
		}

		switch state {

		case HEADER: // beginning of MBOX (thrown away)
			if !bytes.HasPrefix(line, []byte("From ")) {
				return nil, errors.New("MBOX header not found.")
			}
			state = EMPTYLINE
		case EMPTYLINE:
			if bytes.Equal(line, []byte("\r\n")) || bytes.Equal(line, []byte("\n")) {
				eol = line
				state = FOOTERHEADER
			} else {
				if (len(line) == 0) && (err == io.EOF) {
					return nil, errors.New("MBOX footer not found.")
				}

				if escapedFrom.Match(line) {
					line = line[1:]
				}
				msg = append(msg, line...)
			}
		case FOOTERHEADER:
			// Found a header to the next message, so the previous
			// empty line was the footer.
			if err == io.EOF {
				return msg, io.EOF
			}

			if bytes.HasPrefix(line, []byte("From ")) {
				return msg, nil
			}

			// Found another empty line instead header
			// so maybe this is the footer?
			if bytes.Equal(line, []byte("\r\n")) || bytes.Equal(line, []byte("\n")) {
				eol = line
				msg = append(msg, eol...)
				state = FOOTERHEADER
			} else { // found some text, so the empty line was not a footer
				msg = append(msg, eol...)
				eol = nil
				msg = append(msg, line...)
				state = EMPTYLINE
			}
		}
	}
}

func (m *MBox) Next() bool {
	return m.err != io.EOF
}
