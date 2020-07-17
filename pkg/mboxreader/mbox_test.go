package mboxreader

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

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
	mbox := bytes.NewReader([]byte(bytes.Join(raw, []byte(""))))

	m := NewMBOXReader(mbox)

	iterations := 0
	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, raw[1], msg)
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
			`From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Another test
Date: Thu, 03 Jul 2020 22:21:08 -0600

Body of message here
another line

`),
		[]byte("\n"),
		[]byte("From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020\n"),
		[]byte(
			`From: "Some One" <person@example.net>
To: <one@example.com>
Subject: Again
Date: Thu, 04 Jul 2020 22:21:08 -0600

body body body
body body body

`),
		[]byte("\n")}

	mbox := bytes.NewReader([]byte(bytes.Join(msgs, []byte(""))))

	m := NewMBOXReader(mbox)

	offset := 1
	count := 0
	iteration := 0
	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		assert.Equalf(t, msgs[offset], msg, "msg %d", count)

		offset += 3
		count++
		iteration++
	}
	assert.Equal(t, 3, iteration)
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

	mbox := bytes.NewReader(raw)
	m := NewMBOXReader(mbox)

	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

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
	m := NewMBOXReader(bytes.NewReader(errReader(nil)))
	_, err := ioutil.ReadAll(m)
	assert.Error(t, err)
}

func TestBadMBOX(t *testing.T) {
	m := NewMBOXReader(bytes.NewReader([]byte("not the proper header")))
	_, err := ioutil.ReadAll(m)
	assert.EqualError(t, err, "MBOX header not found.")
}

func TestBadMBOX2(t *testing.T) {
	raw := []byte(
		`From 3671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020
From: "Some One" <person@example.net>

foobar
From 9671168074256204903@xxx Fri Jul 05 04:21:10 +0000 2020`)

	mbox := bytes.NewReader(raw)
	m := NewMBOXReader(mbox)
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

	mbox := bytes.NewReader(raw)
	m := NewMBOXReader(mbox)
	msg, err := ioutil.ReadAll(m)
	assert.NoError(t, err)
	assert.Equal(t, want, msg)
	assert.False(t, m.Next())
}
