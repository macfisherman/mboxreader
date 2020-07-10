package mboxreader

import (
	"bytes"
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

	m, err := NewMBOXReader(mbox)
	if err != nil {
		panic(err)
	}

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

	m, err := NewMBOXReader(mbox)
	if err != nil {
		panic(err)
	}

	offset := 1
	count := 0
	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		assert.Equalf(t, msgs[offset], msg, "msg %d", count)

		offset += 3
		count++
	}
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

	m, err := NewMBOXReader(mbox)
	if err != nil {
		panic(err)
	}

	for m.Next() {
		msg, err := ioutil.ReadAll(m)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, want, msg)
	}
}
