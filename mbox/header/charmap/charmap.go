// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright 2021 Jeff Macdonald <macfisherman@gmail.com>

//go:generate go run maketables.go

package charmap

import (
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// stuff from text/encoding/internal
//
// A RepertoireError indicates a rune is not in the repertoire of a destination
// encoding. It is associated with an encoding-specific suggested replacement
// byte.
type RepertoireError byte

// Error implements the error interrface.
func (r RepertoireError) Error() string {
	return "encoding: rune not supported by encoding."
}

// Replacement returns the replacement string associated with this error.
func (r RepertoireError) Replacement() byte { return byte(r) }

var ErrASCIIReplacement = RepertoireError(encoding.ASCIISub)

// All is a list of all defined encodings in this package.
var All []encoding.Encoding = listAll

// utf8Enc holds a rune's UTF-8 encoding in data[:len].
type utf8Enc struct {
	len  uint8
	data [3]byte
}

// Charmap is an 8-bit character set encoding.
type Charmap struct {
	// name is the encoding's name.
	name string
	// mib is the encoding type of this encoder.
	mib int
	// asciiSuperset states whether the encoding is a superset of ASCII.
	asciiSuperset bool
	// low is the lower bound of the encoded byte for a non-ASCII rune. If
	// Charmap.asciiSuperset is true then this will be 0x80, otherwise 0x00.
	low uint8
	// replacement is the encoded replacement character.
	replacement byte
	// decode is the map from encoded byte to UTF-8.
	decode [256]utf8Enc
	// encoding is the map from runes to encoded bytes. Each entry is a
	// uint32: the high 8 bits are the encoded byte and the low 24 bits are
	// the rune. The table entries are sorted by ascending rune.
	encode [256]uint32
}

// NewDecoder implements the encoding.Encoding interface.
func (m *Charmap) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{Transformer: charmapDecoder{charmap: m}}
}

// NewEncoder implements the encoding.Encoding interface.
func (m *Charmap) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{Transformer: charmapEncoder{charmap: m}}
}

// String returns the Charmap's name.
func (m *Charmap) String() string {
	return m.name
}

// ID implements an internal interface.
func (m *Charmap) ID() (mib int, other string) {
	return m.mib, ""
}

// charmapDecoder implements transform.Transformer by decoding to UTF-8.
type charmapDecoder struct {
	transform.NopResetter
	charmap *Charmap
}

func (m charmapDecoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for i, c := range src {
		if m.charmap.asciiSuperset && c < utf8.RuneSelf {
			if nDst >= len(dst) {
				err = transform.ErrShortDst
				break
			}
			dst[nDst] = c
			nDst++
			nSrc = i + 1
			continue
		}

		decode := &m.charmap.decode[c]
		n := int(decode.len)
		if nDst+n > len(dst) {
			err = transform.ErrShortDst
			break
		}
		// It's 15% faster to avoid calling copy for these tiny slices.
		for j := 0; j < n; j++ {
			dst[nDst] = decode.data[j]
			nDst++
		}
		nSrc = i + 1
	}
	return nDst, nSrc, err
}

// DecodeByte returns the Charmap's rune decoding of the byte b.
func (m *Charmap) DecodeByte(b byte) rune {
	switch x := &m.decode[b]; x.len {
	case 1:
		return rune(x.data[0])
	case 2:
		return rune(x.data[0]&0x1f)<<6 | rune(x.data[1]&0x3f)
	default:
		return rune(x.data[0]&0x0f)<<12 | rune(x.data[1]&0x3f)<<6 | rune(x.data[2]&0x3f)
	}
}

// charmapEncoder implements transform.Transformer by encoding from UTF-8.
type charmapEncoder struct {
	transform.NopResetter
	charmap *Charmap
}

func (m charmapEncoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	r, size := rune(0), 0
loop:
	for nSrc < len(src) {
		if nDst >= len(dst) {
			err = transform.ErrShortDst
			break
		}
		r = rune(src[nSrc])

		// Decode a 1-byte rune.
		if r < utf8.RuneSelf {
			if m.charmap.asciiSuperset {
				nSrc++
				dst[nDst] = uint8(r)
				nDst++
				continue
			}
			size = 1

		} else {
			// Decode a multi-byte rune.
			r, size = utf8.DecodeRune(src[nSrc:])
			if size == 1 {
				// All valid runes of size 1 (those below utf8.RuneSelf) were
				// handled above. We have invalid UTF-8 or we haven't seen the
				// full character yet.
				if !atEOF && !utf8.FullRune(src[nSrc:]) {
					err = transform.ErrShortSrc
				} else {
					err = RepertoireError(m.charmap.replacement)
				}
				break
			}
		}

		// Binary search in [low, high) for that rune in the m.charmap.encode table.
		for low, high := int(m.charmap.low), 0x100; ; {
			if low >= high {
				err = RepertoireError(m.charmap.replacement)
				break loop
			}
			mid := (low + high) / 2
			got := m.charmap.encode[mid]
			gotRune := rune(got & (1<<24 - 1))
			if gotRune < r {
				low = mid + 1
			} else if gotRune > r {
				high = mid
			} else {
				dst[nDst] = byte(got >> 24)
				nDst++
				break
			}
		}
		nSrc += size
	}
	return nDst, nSrc, err
}

// EncodeRune returns the Charmap's byte encoding of the rune r. ok is whether
// r is in the Charmap's repertoire. If not, b is set to the Charmap's
// replacement byte. This is often the ASCII substitute character '\x1a'.
func (m *Charmap) EncodeRune(r rune) (b byte, ok bool) {
	if r < utf8.RuneSelf && m.asciiSuperset {
		return byte(r), true
	}
	for low, high := int(m.low), 0x100; ; {
		if low >= high {
			return m.replacement, false
		}
		mid := (low + high) / 2
		got := m.encode[mid]
		gotRune := rune(got & (1<<24 - 1))
		if gotRune < r {
			low = mid + 1
		} else if gotRune > r {
			high = mid
		} else {
			return byte(got >> 24), true
		}
	}
}
