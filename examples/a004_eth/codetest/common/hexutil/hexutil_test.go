package hexutil

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type marshalTest struct {
	input interface{}
	want  string
}

type unmarshalTest struct {
	input        string
	want         interface{}
	wantErr      error // if set, decoding must fail on any platform
	wantErr32bit error // if set, decoding must fail on 32bit platforms (used for Uint tests)
}

// []byte数组转成带0x前缀的16进制字符串
func TestEncode(t *testing.T) {
	encodeBytesTests := []marshalTest{
		{[]byte{}, "0x"},
		{[]byte{0}, "0x00"},
		{[]byte{0, 0, 1, 2}, "0x00000102"},
	}
	for _, test := range encodeBytesTests {
		enc := hexutil.Encode(test.input.([]byte)) // []byte数组转成带0x前缀的16进制字符串
		if enc != test.want {
			t.Errorf("input %x: wrong encoding %s", test.input, enc)
		} else {
			fmt.Println(enc)
			fmt.Println("----")
		}
	}
}

// 16进制字符串转[]byte数组，16进制字符串必须以0x开头，并且是偶数位
func TestDecode(t *testing.T) {
	decodeBytesTests := []unmarshalTest{
		// invalid
		{input: ``, wantErr: hexutil.ErrEmptyString},
		{input: `0`, wantErr: hexutil.ErrMissingPrefix},
		{input: `0x0`, wantErr: hexutil.ErrOddLength},
		{input: `0x023`, wantErr: hexutil.ErrOddLength},
		{input: `0xxx`, wantErr: hexutil.ErrSyntax},
		{input: `0x01zz01`, wantErr: hexutil.ErrSyntax},
		// valid
		{input: `0x`, want: []byte{}},
		{input: `0X`, want: []byte{}},
		{input: `0x02`, want: []byte{0x02}},
		{input: `0X02`, want: []byte{0x02}},
		{input: `0xffffffffff`, want: []byte{0xff, 0xff, 0xff, 0xff, 0xff}},
		{
			input: `0xffffffffffffffffffffffffffffffffffff`,
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}
	for _, test := range decodeBytesTests {
		dec, err := hexutil.Decode(test.input) // 16进制字符串转[]byte数组，16进制字符串必须以0x开头，并且是偶数位
		if !checkError(t, test.input, err, test.wantErr) {
			fmt.Println(test.input, "错误1")
			fmt.Println("----")
			continue
		}
		if !bytes.Equal(test.want.([]byte), dec) {
			fmt.Println(test.input, "错误2")
			fmt.Println("----")
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, dec, test.want)
			continue
		} else {
			fmt.Println(test.input)
			fmt.Println(dec)
			fmt.Println("----")
		}
	}
}

// big.Int转成带0x前缀的16进制字符串
func TestEncodeBig(t *testing.T) {
	encodeBigTests := []marshalTest{
		{referenceBig("0"), "0x0"},
		{referenceBig("1"), "0x1"},
		{referenceBig("ff"), "0xff"},
		{referenceBig("112233445566778899aabbccddeeff"), "0x112233445566778899aabbccddeeff"},
		{referenceBig("80a7f2c1bcc396c00"), "0x80a7f2c1bcc396c00"},
		{referenceBig("-80a7f2c1bcc396c00"), "-0x80a7f2c1bcc396c00"},
	}
	for _, test := range encodeBigTests {
		enc := hexutil.EncodeBig(test.input.(*big.Int)) // big.Int转成带0x前缀的16进制字符串
		if enc != test.want {
			t.Errorf("input %x: wrong encoding %s", test.input, enc)
		} else {
			fmt.Println(test.input)
			fmt.Println(enc)
			fmt.Println("----")
		}
	}
}

// 十六进制字符串转big.Int，十六进制字符串必须以0x开头，不能以0开头，比如0x01是会失败的，可以奇数位。
func TestDecodeBig(t *testing.T) {
	decodeBigTests := []unmarshalTest{
		// invalid
		{input: `0`, wantErr: hexutil.ErrMissingPrefix},
		{input: `0x`, wantErr: hexutil.ErrEmptyNumber},
		{input: `0x01`, wantErr: hexutil.ErrLeadingZero},
		{input: `0xx`, wantErr: hexutil.ErrSyntax},
		{input: `0x1zz01`, wantErr: hexutil.ErrSyntax},
		{
			input:   `0x10000000000000000000000000000000000000000000000000000000000000000`,
			wantErr: hexutil.ErrBig256Range,
		},
		// valid
		{input: `0x0`, want: big.NewInt(0)},
		{input: `0x2`, want: big.NewInt(0x2)},
		{input: `0x2F2`, want: big.NewInt(0x2f2)},
		{input: `0X2F2`, want: big.NewInt(0x2f2)},
		{input: `0x1122aaff`, want: big.NewInt(0x1122aaff)},
		{input: `0xbBb`, want: big.NewInt(0xbbb)},
		{input: `0xfffffffff`, want: big.NewInt(0xfffffffff)},
		{
			input: `0x112233445566778899aabbccddeeff`,
			want:  referenceBig("112233445566778899aabbccddeeff"),
		},
		{
			input: `0xffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("ffffffffffffffffffffffffffffffffffff"),
		},
		{
			input: `0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff`,
			want:  referenceBig("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
		},
	}
	for _, test := range decodeBigTests {
		dec, err := hexutil.DecodeBig(test.input) // 十六进制字符串转big.Int，十六进制字符串必须以0x开头，不能以0开头，比如0x01是会失败的，可以奇数位。
		if !checkError(t, test.input, err, test.wantErr) {
			fmt.Println(test.input, "错误1")
			fmt.Println("----")
			continue
		}
		if dec.Cmp(test.want.(*big.Int)) != 0 {
			fmt.Println(test.input, "错误2")
			fmt.Println("----")
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, dec, test.want)
			continue
		}

		fmt.Println(test.input)
		fmt.Println(dec)
		fmt.Println("----")

	}
}

// uint64转成带0x前缀的16进制字符串,16进制可以奇数位
func TestEncodeUint64(t *testing.T) {
	encodeUint64Tests := []marshalTest{
		{uint64(0), "0x0"},
		{uint64(1), "0x1"},
		{uint64(0xff), "0xff"},
		{uint64(0x1122334455667788), "0x1122334455667788"},
	}
	for _, test := range encodeUint64Tests {
		enc := hexutil.EncodeUint64(test.input.(uint64))
		if enc != test.want {
			t.Errorf("input %x: wrong encoding %s", test.input, enc)
		}
		fmt.Println(test.input)
		fmt.Println(enc)
		fmt.Println("----")
	}
}

// 十六进制字符串转uint64，十六进制字符串必须以0x开头，不能以0开头，比如0x01是会失败的，可以奇数位。
func TestDecodeUint64(t *testing.T) {
	decodeUint64Tests := []unmarshalTest{
		// invalid
		{input: `0`, wantErr: hexutil.ErrMissingPrefix},
		{input: `0x`, wantErr: hexutil.ErrEmptyNumber},
		{input: `0x01`, wantErr: hexutil.ErrLeadingZero},
		{input: `0xfffffffffffffffff`, wantErr: hexutil.ErrUint64Range},
		{input: `0xx`, wantErr: hexutil.ErrSyntax},
		{input: `0x1zz01`, wantErr: hexutil.ErrSyntax},
		// valid
		{input: `0x0`, want: uint64(0)},
		{input: `0x2`, want: uint64(0x2)},
		{input: `0x2F2`, want: uint64(0x2f2)},
		{input: `0X2F2`, want: uint64(0x2f2)},
		{input: `0x1122aaff`, want: uint64(0x1122aaff)},
		{input: `0xbbb`, want: uint64(0xbbb)},
		{input: `0xffffffffffffffff`, want: uint64(0xffffffffffffffff)},
	}
	for _, test := range decodeUint64Tests {
		dec, err := hexutil.DecodeUint64(test.input)
		if !checkError(t, test.input, err, test.wantErr) {
			fmt.Println(test.input, "错误1")
			fmt.Println("----")
			continue
		}
		if dec != test.want.(uint64) {
			fmt.Println(test.input, "错误2")
			fmt.Println("----")
			t.Errorf("input %s: value mismatch: got %x, want %x", test.input, dec, test.want)
			continue
		}
		fmt.Println(test.input)
		fmt.Println(dec)
		fmt.Println("----")
	}
}
