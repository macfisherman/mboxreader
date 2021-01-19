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

package mbox_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/macfisherman/mboxreader/mbox"
	"github.com/stretchr/testify/assert"
)

func TestSingleMessage(t *testing.T) {
	raw := [][]byte{
		[]byte("From 1671168074256204903@xxx Fri Jul 03 04:21:10 +0000 2020\n"),
		[]byte(
			`From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Testing
Date: Thu, 02 Jul 2020 22:21:08 -0600

Body of message here
and here
and here



`),
		[]byte("\n")}

	box := bytes.NewReader([]byte(bytes.Join(raw, []byte(""))))

	m := mbox.New(box)

	iterations := 0
	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		// last one will have entire content till eof
		expected := append(raw[1], byte('\n'))
		assert.Equal(t, expected, msg)
		iterations++
	}

	assert.Equal(t, 1, iterations)
}

func TestMBOX(t *testing.T) {
	msgs := [][]byte{
		[]byte("From 1671168074256204903@xxx Fri Jul 03 04:21:10 +0000 2020\n"),
		[]byte(
			`From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Testing
Date: Thu, 02 Jul 2020 22:21:08 -0600

Body of message here
and here
and here
`),
		[]byte("\n"),
		[]byte("From 2671168074256204903@xxx Fri Jul 04 04:21:10 +0000 2020\n"),
		[]byte(
			`From: "Some Two" <person@example.net>
To: <one@example.com>
Subject: Another test
Date: Thu, 03 Jul 2020 22:21:08 -0600

Body of message here
another line

`),
		[]byte("\n"),
		[]byte("From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020\n"),
		[]byte(
			`From: "Some Three" <person@example.net>
To: <one@example.com>
Subject: Again
Date: Thu, 04 Jul 2020 22:21:08 -0600

body body body
body body body

`),
		[]byte("\n")}

	box := bytes.NewReader([]byte(bytes.Join(msgs, []byte(""))))

	m := mbox.New(box)

	var got [][]byte

	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		got = append(got, msg)
	}

	assert.Equal(t, msgs[1], got[0])
	assert.Equal(t, msgs[4], got[1])

	// last one will have entire content till eof
	expected := append(msgs[7], byte('\n'))
	assert.Equal(t, expected, got[2])
}

func TestEscapedFrom(t *testing.T) {
	raw := []byte(
		`From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020
From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Again
Date: Thu, 04 Jul 2020 22:21:08 -0600

body body body
>From someone
body body body
>>From someone said
>>>From one

`)

	want := []byte(
		`From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Again
Date: Thu, 04 Jul 2020 22:21:08 -0600

body body body
From someone
body body body
>From someone said
>>From one

`)

	box := bytes.NewReader(raw)
	m := mbox.New(box)

	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		t.Log(msg)

		assert.Equal(t, want, msg)
	}
}

type errReader []byte

func (errReader) Read(p []byte) (int, error) {
	return 0, errors.New("test read error")
}

type errShortReader []byte

func (e errShortReader) Read(p []byte) (int, error) {
	n := copy(p, e)
	return n, errors.New("short read")
}

func TestReadErrors(t *testing.T) {
	m := mbox.New(bytes.NewReader(errReader(nil)))
	_, err := ioutil.ReadAll(m)
	assert.Error(t, err)
}

func TestBadMBOX(t *testing.T) {
	m := mbox.New(bytes.NewReader([]byte("not the proper header")))
	_, err := ioutil.ReadAll(m)
	assert.EqualError(t, err, "MBOX header not found.")
}

func TestBadMBOX2(t *testing.T) {
	raw := []byte(
		`From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020
From: "Some One" <person@example.net>

foobar
From 9671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020`)

	box := bytes.NewReader(raw)
	m := mbox.New(box)
	_, err := ioutil.ReadAll(m)
	assert.EqualError(t, err, "MBOX footer not found.")
}

func TestBadMBOX3(t *testing.T) {
	raw := []byte(
		`From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020
From: "Some One" <person@example.net>

foobar

From 9671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020`)

	want := []byte(`From: "Some One" <person@example.net>

foobar
`)

	box := bytes.NewReader(raw)
	m := mbox.New(box)
	msg, err := ioutil.ReadAll(m)
	assert.NoError(t, err)
	assert.Equal(t, want, msg)
	assert.False(t, m.Next())
}
