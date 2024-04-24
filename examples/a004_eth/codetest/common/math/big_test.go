// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package math

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// 数字字符串转math.HexOrDecimal256,这是big.Int
func TestHexOrDecimal256(t *testing.T) {
	tests := []struct {
		input string
		num   *big.Int
		ok    bool
	}{
		{"", big.NewInt(0), true},
		{"0", big.NewInt(0), true},
		{"0x0", big.NewInt(0), true},
		{"12345678", big.NewInt(12345678), true},
		{"0x12345678", big.NewInt(0x12345678), true},
		{"0X12345678", big.NewInt(0x12345678), true},
		// Tests for leading zero behaviour:
		{"0123456789", big.NewInt(123456789), true}, // note: not octal
		{"00", big.NewInt(0), true},
		{"0x00", big.NewInt(0), true},
		{"0x012345678abc", big.NewInt(0x12345678abc), true},
		// Invalid syntax:
		{"abcdef", nil, false},
		{"0xgg", nil, false},
		// Larger than 256 bits:
		{"115792089237316195423570985008687907853269984665640564039457584007913129639936", nil, false},
	}
	for _, test := range tests {
		var num math.HexOrDecimal256
		err := num.UnmarshalText([]byte(test.input))
		if (err == nil) != test.ok {
			t.Errorf("ParseBig(%q) -> (err == nil) == %t, want %t", test.input, err == nil, test.ok)
			continue
		}
		if test.num != nil && (*big.Int)(&num).Cmp(test.num) != 0 {
			t.Errorf("ParseBig(%q) -> %d, want %d", test.input, (*big.Int)(&num), test.num)
		}
		fmt.Println(test.input)
		fmt.Println(num)
		fmt.Println("----")
	}
}

// 数字字符串转math.HexOrDecimal256,这是big.Int
func TestMustParseBig256(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MustParseBig should've panicked")
		} else {
			fmt.Println("转换失败")
		}
	}()
	r := math.MustParseBig256("ggg")
	fmt.Println("r = ", r)
}

// 最大值
func TestBigMax(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(5)

	max1 := math.BigMax(a, b)
	if max1 != a {
		t.Errorf("Expected %d got %d", a, max1)
	}

	max2 := math.BigMax(b, a)
	if max2 != a {
		t.Errorf("Expected %d got %d", a, max2)
	}
}

// 最小值
func TestBigMin(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(5)

	min1 := math.BigMin(a, b)
	if min1 != b {
		t.Errorf("Expected %d got %d", b, min1)
	}

	min2 := math.BigMin(b, a)
	if min2 != b {
		t.Errorf("Expected %d got %d", b, min2)
	}
}

// 最右1的索引，从右往左，从0开始
func TestFirstBigSet(t *testing.T) {
	tests := []struct {
		num *big.Int
		ix  int
	}{
		{big.NewInt(0), 0},
		{big.NewInt(1), 0},
		{big.NewInt(6), 1},
		{big.NewInt(0x100), 8},
	}
	for _, test := range tests {
		fmt.Println(math.FirstBitSet(test.num)) // 最右1的索引，从右往左，从0开始
		continue
		if ix := math.FirstBitSet(test.num); ix != test.ix {
			t.Errorf("FirstBitSet(b%b) = %d, want %d", test.num, ix, test.ix)
		}
	}
}

// big.Int转大端序字节数组
func TestPaddedBigBytes(t *testing.T) {
	tests := []struct {
		num    *big.Int
		n      int
		result []byte
	}{
		{num: big.NewInt(0), n: 4, result: []byte{0, 0, 0, 0}},
		{num: big.NewInt(3), n: 5, result: []byte{0, 0, 0, 1}},
		{num: big.NewInt(256), n: 4, result: []byte{0, 0, 2, 0}},
		{num: math.BigPow(2, 32), n: 4, result: []byte{1, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		// if result := math.PaddedBigBytes(test.num, test.n); !bytes.Equal(result, test.result) {
		// 	t.Errorf("PaddedBigBytes(%d, %d) = %v, want %v", test.num, test.n, result, test.result)
		// }
		fmt.Println(test.num, test.n)
		fmt.Println(math.PaddedBigBytes(test.num, test.n)) // big.Int转大端序字节数组
		fmt.Println("----")
	}
}

// bigEndianByteAt returns the byte at position n,
// in Big-Endian encoding
// So n==0 returns the least significant byte
func bigEndianByteAt(bigint *big.Int, n int) byte {
	words := bigint.Bits()
	// Check word-bucket the byte will reside in
	i := n / wordBytes
	if i >= len(words) {
		return byte(0)
	}
	word := words[i]
	// Offset of the byte
	shift := 8 * uint(n%wordBytes)

	return byte(word >> shift)
}

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// 这段代码定义了一个名为 ReadBits 的函数，它用于将 big.Int 类型的大整数的绝对值编码为大端字节切片。函数的调用者需要确保提供的 buf 字节切片有足够的空间来存储编码后的字节。如果 buf 太短，则结果将是不完整的。
// 将big.Int转换成大端序的字节切片
func TestReadBits(t *testing.T) {
	check := func(input string) {
		want, _ := hex.DecodeString(input)
		int, _ := new(big.Int).SetString(input, 16)
		buf := make([]byte, len(want))
		math.ReadBits(int, buf) //将big.Int转换成大端序的字节切片
		fmt.Println(int.Text(16))
		fmt.Println(buf, want)
		if !bytes.Equal(buf, want) {
			t.Errorf("have: %x\nwant: %x", buf, want)
		}
	}
	check("1234")
	// check("000000000000000000000000000000000000000000000000000000FEFCF3F8F0")
	// check("0000000000012345000000000000000000000000000000000000FEFCF3F8F0")
	// check("18F8F8F1000111000110011100222004330052300000000000000000FEFCF3F8F0")
}

// u256类型，-1也能转换成u256
func TestU256(t *testing.T) {
	tests := []struct{ x, y *big.Int }{
		{x: big.NewInt(0), y: big.NewInt(0)},
		{x: big.NewInt(1), y: big.NewInt(1)},
		{x: math.BigPow(2, 255), y: math.BigPow(2, 255)},
		{x: math.BigPow(2, 256), y: big.NewInt(0)},
		{x: new(big.Int).Add(math.BigPow(2, 256), big.NewInt(1)), y: big.NewInt(1)},
		// negative values
		{x: big.NewInt(-1), y: new(big.Int).Sub(math.BigPow(2, 256), big.NewInt(1))},
		{x: big.NewInt(-2), y: new(big.Int).Sub(math.BigPow(2, 256), big.NewInt(2))},
		{x: math.BigPow(2, -255), y: big.NewInt(1)},
	}
	for _, test := range tests {
		if y := math.U256(new(big.Int).Set(test.x)); y.Cmp(test.y) != 0 { //big.Int转换成u256类型
			t.Errorf("U256(%x) = %x, want %x", test.x, y, test.y)
		}
		fmt.Println(test.x, math.U256(new(big.Int).Set(test.x)))
		fmt.Println("----")
	}
}

// u256类型转换成[]byte类型，-1也能转换成[]byte，大端模式
func TestU256Bytes(t *testing.T) {
	ubytes := make([]byte, 32)
	ubytes[31] = 1

	unsigned := math.U256Bytes(big.NewInt(-1)) //u256类型转换成[]byte类型，-1也能转换成[]byte，大端模式
	// if !bytes.Equal(unsigned, ubytes) {
	// 	t.Errorf("expected %x got %x", ubytes, unsigned)
	// }
	fmt.Println(unsigned)
}

// 大端模式
func TestBigEndianByteAt(t *testing.T) {

	tests := []struct {
		x   string
		y   int
		exp byte
	}{
		{"00", 0, 0x00},
		{"01", 1, 0x00},
		{"00", 1, 0x00},
		{"01", 0, 0x01},
		{"0000000000000000000000000000000000000000000000000000000000102030", 0, 0x30},
		{"0000000000000000000000000000000000000000000000000000000000102030", 1, 0x20},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 31, 0xAB},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 32, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 500, 0x00},
	}
	for _, test := range tests {
		v := new(big.Int).SetBytes(common.Hex2Bytes(test.x))
		actual := bigEndianByteAt(v, test.y)
		if actual != test.exp {
			t.Fatalf("Expected  [%v] %v:th byte to be %v, was %v.", test.x, test.y, test.exp, actual)
		}
		fmt.Println(test.x, v.Text(16), test.y, test.exp, actual)
		fmt.Println("----")
	}
}

// 小端模式
func TestLittleEndianByteAt(t *testing.T) {
	// fmt.Println(big.NewInt(258).Bytes())
	// return
	tests := []struct {
		x   string
		y   int
		exp byte
	}{
		{"00", 0, 0x00},
		{"01", 1, 0x00},
		{"00", 1, 0x00},
		{"01", 0, 0x00},
		{"0000000000000000000000000000000000000000000000000000000000102030", 0, 0x00},
		{"0000000000000000000000000000000000000000000000000000000000102030", 1, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 31, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 32, 0x00},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 0, 0xAB},
		{"ABCDEF0908070605040302010000000000000000000000000000000000000000", 1, 0xCD},
		{"00CDEF090807060504030201ffffffffffffffffffffffffffffffffffffffff", 0, 0x00},
		{"00CDEF090807060504030201ffffffffffffffffffffffffffffffffffffffff", 1, 0xCD},
		{"0000000000000000000000000000000000000000000000000000000000102030", 31, 0x30},
		{"0000000000000000000000000000000000000000000000000000000000102030", 30, 0x20},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 32, 0x0},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 31, 0xFF},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0xFFFF, 0x0},
	}
	for _, test := range tests {
		v := new(big.Int).SetBytes(common.Hex2Bytes(test.x))
		actual := math.Byte(v, 32, test.y)
		if actual != test.exp {
			t.Fatalf("Expected  [%v] %v:th byte to be %v, was %v.", test.x, test.y, test.exp, actual)
		}
		fmt.Println(test.x, v.Text(16), test.y, test.exp, actual)
		fmt.Println("----")
	}
}

//有符号256位数

func TestS256(t *testing.T) {
	tests := []struct{ x, y *big.Int }{
		{x: big.NewInt(0), y: big.NewInt(0)},
		{x: big.NewInt(1), y: big.NewInt(1)},
		{x: big.NewInt(2), y: big.NewInt(2)},
		{
			x: new(big.Int).Sub(math.BigPow(2, 255), big.NewInt(1)),
			y: new(big.Int).Sub(math.BigPow(2, 255), big.NewInt(1)),
		},
		{
			x: math.BigPow(2, 255),
			y: new(big.Int).Neg(math.BigPow(2, 255)),
		},
		{
			x: new(big.Int).Sub(math.BigPow(2, 256), big.NewInt(1)),
			y: big.NewInt(-1),
		},
		{
			x: new(big.Int).Sub(math.BigPow(2, 256), big.NewInt(2)),
			y: big.NewInt(-2),
		},
	}
	for _, test := range tests {
		if y := math.S256(test.x); y.Cmp(test.y) != 0 {
			t.Errorf("S256(%x) = %x, want %x", test.x, y, test.y)
		}
		fmt.Println(test.x, test.y, math.S256(test.x))
		fmt.Println("----")
	}
}

// a的b次方
func TestExp(t *testing.T) {
	tests := []struct{ base, exponent, result *big.Int }{
		{base: big.NewInt(0), exponent: big.NewInt(0), result: big.NewInt(1)},
		{base: big.NewInt(3), exponent: big.NewInt(3), result: big.NewInt(1)},
		{base: big.NewInt(1), exponent: big.NewInt(1), result: big.NewInt(1)},
		{base: big.NewInt(1), exponent: big.NewInt(2), result: big.NewInt(1)},
		{base: big.NewInt(3), exponent: big.NewInt(144), result: math.MustParseBig256("507528786056415600719754159741696356908742250191663887263627442114881")},
		{base: big.NewInt(2), exponent: big.NewInt(255), result: math.MustParseBig256("57896044618658097711785492504343953926634992332820282019728792003956564819968")},
	}
	for _, test := range tests {
		// if result := math.Exp(test.base, test.exponent); result.Cmp(test.result) != 0 {
		// 	t.Errorf("Exp(%d, %d) = %d, want %d", test.base, test.exponent, result, test.result)
		// }
		fmt.Println(test.base, test.exponent, test.result.Text(16), math.Exp(test.base, test.exponent).Text(16))
		fmt.Println("----")
	}
}
