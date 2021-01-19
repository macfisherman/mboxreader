// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.
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

package charmap

import "golang.org/x/text/encoding"

// hp-roman8 is the HP Roman8 encoding.
var HPRoman8 *Charmap = &hproman8

var hproman8 = Charmap{
	name:          "HP Roman8",
	mib:           2004,
	asciiSuperset: true,
	low:           0x80,
	replacement:   0x1a,
	decode: [256]utf8Enc{
		{1, [3]byte{0x00, 0x00, 0x00}}, {1, [3]byte{0x01, 0x00, 0x00}},
		{1, [3]byte{0x02, 0x00, 0x00}}, {1, [3]byte{0x03, 0x00, 0x00}},
		{1, [3]byte{0x04, 0x00, 0x00}}, {1, [3]byte{0x05, 0x00, 0x00}},
		{1, [3]byte{0x06, 0x00, 0x00}}, {1, [3]byte{0x07, 0x00, 0x00}},
		{1, [3]byte{0x08, 0x00, 0x00}}, {1, [3]byte{0x09, 0x00, 0x00}},
		{1, [3]byte{0x0a, 0x00, 0x00}}, {1, [3]byte{0x0b, 0x00, 0x00}},
		{1, [3]byte{0x0c, 0x00, 0x00}}, {1, [3]byte{0x0d, 0x00, 0x00}},
		{1, [3]byte{0x0e, 0x00, 0x00}}, {1, [3]byte{0x0f, 0x00, 0x00}},
		{1, [3]byte{0x10, 0x00, 0x00}}, {1, [3]byte{0x11, 0x00, 0x00}},
		{1, [3]byte{0x12, 0x00, 0x00}}, {1, [3]byte{0x13, 0x00, 0x00}},
		{1, [3]byte{0x14, 0x00, 0x00}}, {1, [3]byte{0x15, 0x00, 0x00}},
		{1, [3]byte{0x16, 0x00, 0x00}}, {1, [3]byte{0x17, 0x00, 0x00}},
		{1, [3]byte{0x18, 0x00, 0x00}}, {1, [3]byte{0x19, 0x00, 0x00}},
		{1, [3]byte{0x1a, 0x00, 0x00}}, {1, [3]byte{0x1b, 0x00, 0x00}},
		{1, [3]byte{0x1c, 0x00, 0x00}}, {1, [3]byte{0x1d, 0x00, 0x00}},
		{1, [3]byte{0x1e, 0x00, 0x00}}, {1, [3]byte{0x1f, 0x00, 0x00}},
		{1, [3]byte{0x20, 0x00, 0x00}}, {1, [3]byte{0x21, 0x00, 0x00}},
		{1, [3]byte{0x22, 0x00, 0x00}}, {1, [3]byte{0x23, 0x00, 0x00}},
		{1, [3]byte{0x24, 0x00, 0x00}}, {1, [3]byte{0x25, 0x00, 0x00}},
		{1, [3]byte{0x26, 0x00, 0x00}}, {1, [3]byte{0x27, 0x00, 0x00}},
		{1, [3]byte{0x28, 0x00, 0x00}}, {1, [3]byte{0x29, 0x00, 0x00}},
		{1, [3]byte{0x2a, 0x00, 0x00}}, {1, [3]byte{0x2b, 0x00, 0x00}},
		{1, [3]byte{0x2c, 0x00, 0x00}}, {1, [3]byte{0x2d, 0x00, 0x00}},
		{1, [3]byte{0x2e, 0x00, 0x00}}, {1, [3]byte{0x2f, 0x00, 0x00}},
		{1, [3]byte{0x30, 0x00, 0x00}}, {1, [3]byte{0x31, 0x00, 0x00}},
		{1, [3]byte{0x32, 0x00, 0x00}}, {1, [3]byte{0x33, 0x00, 0x00}},
		{1, [3]byte{0x34, 0x00, 0x00}}, {1, [3]byte{0x35, 0x00, 0x00}},
		{1, [3]byte{0x36, 0x00, 0x00}}, {1, [3]byte{0x37, 0x00, 0x00}},
		{1, [3]byte{0x38, 0x00, 0x00}}, {1, [3]byte{0x39, 0x00, 0x00}},
		{1, [3]byte{0x3a, 0x00, 0x00}}, {1, [3]byte{0x3b, 0x00, 0x00}},
		{1, [3]byte{0x3c, 0x00, 0x00}}, {1, [3]byte{0x3d, 0x00, 0x00}},
		{1, [3]byte{0x3e, 0x00, 0x00}}, {1, [3]byte{0x3f, 0x00, 0x00}},
		{1, [3]byte{0x40, 0x00, 0x00}}, {1, [3]byte{0x41, 0x00, 0x00}},
		{1, [3]byte{0x42, 0x00, 0x00}}, {1, [3]byte{0x43, 0x00, 0x00}},
		{1, [3]byte{0x44, 0x00, 0x00}}, {1, [3]byte{0x45, 0x00, 0x00}},
		{1, [3]byte{0x46, 0x00, 0x00}}, {1, [3]byte{0x47, 0x00, 0x00}},
		{1, [3]byte{0x48, 0x00, 0x00}}, {1, [3]byte{0x49, 0x00, 0x00}},
		{1, [3]byte{0x4a, 0x00, 0x00}}, {1, [3]byte{0x4b, 0x00, 0x00}},
		{1, [3]byte{0x4c, 0x00, 0x00}}, {1, [3]byte{0x4d, 0x00, 0x00}},
		{1, [3]byte{0x4e, 0x00, 0x00}}, {1, [3]byte{0x4f, 0x00, 0x00}},
		{1, [3]byte{0x50, 0x00, 0x00}}, {1, [3]byte{0x51, 0x00, 0x00}},
		{1, [3]byte{0x52, 0x00, 0x00}}, {1, [3]byte{0x53, 0x00, 0x00}},
		{1, [3]byte{0x54, 0x00, 0x00}}, {1, [3]byte{0x55, 0x00, 0x00}},
		{1, [3]byte{0x56, 0x00, 0x00}}, {1, [3]byte{0x57, 0x00, 0x00}},
		{1, [3]byte{0x58, 0x00, 0x00}}, {1, [3]byte{0x59, 0x00, 0x00}},
		{1, [3]byte{0x5a, 0x00, 0x00}}, {1, [3]byte{0x5b, 0x00, 0x00}},
		{1, [3]byte{0x5c, 0x00, 0x00}}, {1, [3]byte{0x5d, 0x00, 0x00}},
		{1, [3]byte{0x5e, 0x00, 0x00}}, {1, [3]byte{0x5f, 0x00, 0x00}},
		{1, [3]byte{0x60, 0x00, 0x00}}, {1, [3]byte{0x61, 0x00, 0x00}},
		{1, [3]byte{0x62, 0x00, 0x00}}, {1, [3]byte{0x63, 0x00, 0x00}},
		{1, [3]byte{0x64, 0x00, 0x00}}, {1, [3]byte{0x65, 0x00, 0x00}},
		{1, [3]byte{0x66, 0x00, 0x00}}, {1, [3]byte{0x67, 0x00, 0x00}},
		{1, [3]byte{0x68, 0x00, 0x00}}, {1, [3]byte{0x69, 0x00, 0x00}},
		{1, [3]byte{0x6a, 0x00, 0x00}}, {1, [3]byte{0x6b, 0x00, 0x00}},
		{1, [3]byte{0x6c, 0x00, 0x00}}, {1, [3]byte{0x6d, 0x00, 0x00}},
		{1, [3]byte{0x6e, 0x00, 0x00}}, {1, [3]byte{0x6f, 0x00, 0x00}},
		{1, [3]byte{0x70, 0x00, 0x00}}, {1, [3]byte{0x71, 0x00, 0x00}},
		{1, [3]byte{0x72, 0x00, 0x00}}, {1, [3]byte{0x73, 0x00, 0x00}},
		{1, [3]byte{0x74, 0x00, 0x00}}, {1, [3]byte{0x75, 0x00, 0x00}},
		{1, [3]byte{0x76, 0x00, 0x00}}, {1, [3]byte{0x77, 0x00, 0x00}},
		{1, [3]byte{0x78, 0x00, 0x00}}, {1, [3]byte{0x79, 0x00, 0x00}},
		{1, [3]byte{0x7a, 0x00, 0x00}}, {1, [3]byte{0x7b, 0x00, 0x00}},
		{1, [3]byte{0x7c, 0x00, 0x00}}, {1, [3]byte{0x7d, 0x00, 0x00}},
		{1, [3]byte{0x7e, 0x00, 0x00}}, {1, [3]byte{0x7f, 0x00, 0x00}},
		{2, [3]byte{0xc2, 0x80, 0x00}}, {2, [3]byte{0xc2, 0x81, 0x00}},
		{2, [3]byte{0xc2, 0x82, 0x00}}, {2, [3]byte{0xc2, 0x83, 0x00}},
		{2, [3]byte{0xc2, 0x84, 0x00}}, {2, [3]byte{0xc2, 0x85, 0x00}},
		{2, [3]byte{0xc2, 0x86, 0x00}}, {2, [3]byte{0xc2, 0x87, 0x00}},
		{2, [3]byte{0xc2, 0x88, 0x00}}, {2, [3]byte{0xc2, 0x89, 0x00}},
		{2, [3]byte{0xc2, 0x8a, 0x00}}, {2, [3]byte{0xc2, 0x8b, 0x00}},
		{2, [3]byte{0xc2, 0x8c, 0x00}}, {2, [3]byte{0xc2, 0x8d, 0x00}},
		{2, [3]byte{0xc2, 0x8e, 0x00}}, {2, [3]byte{0xc2, 0x8f, 0x00}},
		{2, [3]byte{0xc2, 0x90, 0x00}}, {2, [3]byte{0xc2, 0x91, 0x00}},
		{2, [3]byte{0xc2, 0x92, 0x00}}, {2, [3]byte{0xc2, 0x93, 0x00}},
		{2, [3]byte{0xc2, 0x94, 0x00}}, {2, [3]byte{0xc2, 0x95, 0x00}},
		{2, [3]byte{0xc2, 0x96, 0x00}}, {2, [3]byte{0xc2, 0x97, 0x00}},
		{2, [3]byte{0xc2, 0x98, 0x00}}, {2, [3]byte{0xc2, 0x99, 0x00}},
		{2, [3]byte{0xc2, 0x9a, 0x00}}, {2, [3]byte{0xc2, 0x9b, 0x00}},
		{2, [3]byte{0xc2, 0x9c, 0x00}}, {2, [3]byte{0xc2, 0x9d, 0x00}},
		{2, [3]byte{0xc2, 0x9e, 0x00}}, {2, [3]byte{0xc2, 0x9f, 0x00}},
		{2, [3]byte{0xc2, 0xa0, 0x00}}, {2, [3]byte{0xc3, 0x80, 0x00}},
		{2, [3]byte{0xc3, 0x82, 0x00}}, {2, [3]byte{0xc3, 0x88, 0x00}},
		{2, [3]byte{0xc3, 0x8a, 0x00}}, {2, [3]byte{0xc3, 0x8b, 0x00}},
		{2, [3]byte{0xc3, 0x8e, 0x00}}, {2, [3]byte{0xc3, 0x8f, 0x00}},
		{2, [3]byte{0xc2, 0xb4, 0x00}}, {2, [3]byte{0xcb, 0x8b, 0x00}},
		{2, [3]byte{0xcb, 0x86, 0x00}}, {2, [3]byte{0xc2, 0xa8, 0x00}},
		{2, [3]byte{0xcb, 0x9c, 0x00}}, {2, [3]byte{0xc3, 0x99, 0x00}},
		{2, [3]byte{0xc3, 0x9b, 0x00}}, {3, [3]byte{0xe2, 0x82, 0xa4}},
		{2, [3]byte{0xc2, 0xaf, 0x00}}, {2, [3]byte{0xc3, 0x9d, 0x00}},
		{2, [3]byte{0xc3, 0xbd, 0x00}}, {2, [3]byte{0xc2, 0xb0, 0x00}},
		{2, [3]byte{0xc3, 0x87, 0x00}}, {2, [3]byte{0xc3, 0xa7, 0x00}},
		{2, [3]byte{0xc3, 0x91, 0x00}}, {2, [3]byte{0xc3, 0xb1, 0x00}},
		{2, [3]byte{0xc2, 0xa1, 0x00}}, {2, [3]byte{0xc2, 0xbf, 0x00}},
		{2, [3]byte{0xc2, 0xa4, 0x00}}, {2, [3]byte{0xc2, 0xa3, 0x00}},
		{2, [3]byte{0xc2, 0xa5, 0x00}}, {2, [3]byte{0xc2, 0xa7, 0x00}},
		{2, [3]byte{0xc6, 0x92, 0x00}}, {2, [3]byte{0xc2, 0xa2, 0x00}},
		{2, [3]byte{0xc3, 0xa2, 0x00}}, {2, [3]byte{0xc3, 0xaa, 0x00}},
		{2, [3]byte{0xc3, 0xb4, 0x00}}, {2, [3]byte{0xc3, 0xbb, 0x00}},
		{2, [3]byte{0xc3, 0xa1, 0x00}}, {2, [3]byte{0xc3, 0xa9, 0x00}},
		{2, [3]byte{0xc3, 0xb3, 0x00}}, {2, [3]byte{0xc3, 0xba, 0x00}},
		{2, [3]byte{0xc3, 0xa0, 0x00}}, {2, [3]byte{0xc3, 0xa8, 0x00}},
		{2, [3]byte{0xc3, 0xb2, 0x00}}, {2, [3]byte{0xc3, 0xb9, 0x00}},
		{2, [3]byte{0xc3, 0xa4, 0x00}}, {2, [3]byte{0xc3, 0xab, 0x00}},
		{2, [3]byte{0xc3, 0xb6, 0x00}}, {2, [3]byte{0xc3, 0xbc, 0x00}},
		{2, [3]byte{0xc3, 0x85, 0x00}}, {2, [3]byte{0xc3, 0xae, 0x00}},
		{2, [3]byte{0xc3, 0x98, 0x00}}, {2, [3]byte{0xc3, 0x86, 0x00}},
		{2, [3]byte{0xc3, 0xa5, 0x00}}, {2, [3]byte{0xc3, 0xad, 0x00}},
		{2, [3]byte{0xc3, 0xb8, 0x00}}, {2, [3]byte{0xc3, 0xa6, 0x00}},
		{2, [3]byte{0xc3, 0x84, 0x00}}, {2, [3]byte{0xc3, 0xac, 0x00}},
		{2, [3]byte{0xc3, 0x96, 0x00}}, {2, [3]byte{0xc3, 0x9c, 0x00}},
		{2, [3]byte{0xc3, 0x89, 0x00}}, {2, [3]byte{0xc3, 0xaf, 0x00}},
		{2, [3]byte{0xc3, 0x9f, 0x00}}, {2, [3]byte{0xc3, 0x94, 0x00}},
		{2, [3]byte{0xc3, 0x81, 0x00}}, {2, [3]byte{0xc3, 0x83, 0x00}},
		{2, [3]byte{0xc3, 0xa3, 0x00}}, {2, [3]byte{0xc3, 0x90, 0x00}},
		{2, [3]byte{0xc3, 0xb0, 0x00}}, {2, [3]byte{0xc3, 0x8d, 0x00}},
		{2, [3]byte{0xc3, 0x8c, 0x00}}, {2, [3]byte{0xc3, 0x93, 0x00}},
		{2, [3]byte{0xc3, 0x92, 0x00}}, {2, [3]byte{0xc3, 0x95, 0x00}},
		{2, [3]byte{0xc3, 0xb5, 0x00}}, {2, [3]byte{0xc5, 0xa0, 0x00}},
		{2, [3]byte{0xc5, 0xa1, 0x00}}, {2, [3]byte{0xc3, 0x9a, 0x00}},
		{2, [3]byte{0xc5, 0xb8, 0x00}}, {2, [3]byte{0xc3, 0xbf, 0x00}},
		{2, [3]byte{0xc3, 0x9e, 0x00}}, {2, [3]byte{0xc3, 0xbe, 0x00}},
		{2, [3]byte{0xc2, 0xb7, 0x00}}, {2, [3]byte{0xc2, 0xb5, 0x00}},
		{2, [3]byte{0xc2, 0xb6, 0x00}}, {2, [3]byte{0xc2, 0xbe, 0x00}},
		{3, [3]byte{0xe2, 0x80, 0x94}}, {2, [3]byte{0xc2, 0xbc, 0x00}},
		{2, [3]byte{0xc2, 0xbd, 0x00}}, {2, [3]byte{0xc2, 0xaa, 0x00}},
		{2, [3]byte{0xc2, 0xba, 0x00}}, {2, [3]byte{0xc2, 0xab, 0x00}},
		{3, [3]byte{0xe2, 0x96, 0xa0}}, {2, [3]byte{0xc2, 0xbb, 0x00}},
		{2, [3]byte{0xc2, 0xb1, 0x00}}, {3, [3]byte{0xef, 0xbf, 0xbd}},
	},
	encode: [256]uint32{
		0x00000000, 0x01000001, 0x02000002, 0x03000003, 0x04000004, 0x05000005, 0x06000006, 0x07000007,
		0x08000008, 0x09000009, 0x0a00000a, 0x0b00000b, 0x0c00000c, 0x0d00000d, 0x0e00000e, 0x0f00000f,
		0x10000010, 0x11000011, 0x12000012, 0x13000013, 0x14000014, 0x15000015, 0x16000016, 0x17000017,
		0x18000018, 0x19000019, 0x1a00001a, 0x1b00001b, 0x1c00001c, 0x1d00001d, 0x1e00001e, 0x1f00001f,
		0x20000020, 0x21000021, 0x22000022, 0x23000023, 0x24000024, 0x25000025, 0x26000026, 0x27000027,
		0x28000028, 0x29000029, 0x2a00002a, 0x2b00002b, 0x2c00002c, 0x2d00002d, 0x2e00002e, 0x2f00002f,
		0x30000030, 0x31000031, 0x32000032, 0x33000033, 0x34000034, 0x35000035, 0x36000036, 0x37000037,
		0x38000038, 0x39000039, 0x3a00003a, 0x3b00003b, 0x3c00003c, 0x3d00003d, 0x3e00003e, 0x3f00003f,
		0x40000040, 0x41000041, 0x42000042, 0x43000043, 0x44000044, 0x45000045, 0x46000046, 0x47000047,
		0x48000048, 0x49000049, 0x4a00004a, 0x4b00004b, 0x4c00004c, 0x4d00004d, 0x4e00004e, 0x4f00004f,
		0x50000050, 0x51000051, 0x52000052, 0x53000053, 0x54000054, 0x55000055, 0x56000056, 0x57000057,
		0x58000058, 0x59000059, 0x5a00005a, 0x5b00005b, 0x5c00005c, 0x5d00005d, 0x5e00005e, 0x5f00005f,
		0x60000060, 0x61000061, 0x62000062, 0x63000063, 0x64000064, 0x65000065, 0x66000066, 0x67000067,
		0x68000068, 0x69000069, 0x6a00006a, 0x6b00006b, 0x6c00006c, 0x6d00006d, 0x6e00006e, 0x6f00006f,
		0x70000070, 0x71000071, 0x72000072, 0x73000073, 0x74000074, 0x75000075, 0x76000076, 0x77000077,
		0x78000078, 0x79000079, 0x7a00007a, 0x7b00007b, 0x7c00007c, 0x7d00007d, 0x7e00007e, 0x7f00007f,
		0x80000080, 0x81000081, 0x82000082, 0x83000083, 0x84000084, 0x85000085, 0x86000086, 0x87000087,
		0x88000088, 0x89000089, 0x8a00008a, 0x8b00008b, 0x8c00008c, 0x8d00008d, 0x8e00008e, 0x8f00008f,
		0x90000090, 0x91000091, 0x92000092, 0x93000093, 0x94000094, 0x95000095, 0x96000096, 0x97000097,
		0x98000098, 0x99000099, 0x9a00009a, 0x9b00009b, 0x9c00009c, 0x9d00009d, 0x9e00009e, 0x9f00009f,
		0xa00000a0, 0xb80000a1, 0xbf0000a2, 0xbb0000a3, 0xba0000a4, 0xbc0000a5, 0xbd0000a7, 0xab0000a8,
		0xf90000aa, 0xfb0000ab, 0xb00000af, 0xb30000b0, 0xfe0000b1, 0xa80000b4, 0xf30000b5, 0xf40000b6,
		0xf20000b7, 0xfa0000ba, 0xfd0000bb, 0xf70000bc, 0xf80000bd, 0xf50000be, 0xb90000bf, 0xa10000c0,
		0xe00000c1, 0xa20000c2, 0xe10000c3, 0xd80000c4, 0xd00000c5, 0xd30000c6, 0xb40000c7, 0xa30000c8,
		0xdc0000c9, 0xa40000ca, 0xa50000cb, 0xe60000cc, 0xe50000cd, 0xa60000ce, 0xa70000cf, 0xe30000d0,
		0xb60000d1, 0xe80000d2, 0xe70000d3, 0xdf0000d4, 0xe90000d5, 0xda0000d6, 0xd20000d8, 0xad0000d9,
		0xed0000da, 0xae0000db, 0xdb0000dc, 0xb10000dd, 0xf00000de, 0xde0000df, 0xc80000e0, 0xc40000e1,
		0xc00000e2, 0xe20000e3, 0xcc0000e4, 0xd40000e5, 0xd70000e6, 0xb50000e7, 0xc90000e8, 0xc50000e9,
		0xc10000ea, 0xcd0000eb, 0xd90000ec, 0xd50000ed, 0xd10000ee, 0xdd0000ef, 0xe40000f0, 0xb70000f1,
		0xca0000f2, 0xc60000f3, 0xc20000f4, 0xea0000f5, 0xce0000f6, 0xd60000f8, 0xcb0000f9, 0xc70000fa,
		0xc30000fb, 0xcf0000fc, 0xb20000fd, 0xf10000fe, 0xef0000ff, 0xeb000160, 0xec000161, 0xee000178,
		0xbe000192, 0xaa0002c6, 0xa90002cb, 0xac0002dc, 0xf6002014, 0xaf0020a4, 0xfc0025a0, 0xfc0025a0,
	},
}
var listAll = []encoding.Encoding{
	HPRoman8,
}

// Total table size 2072 bytes (2KiB); checksum: 811C9DC5