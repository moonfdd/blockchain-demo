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
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
)

type operation byte

const (
	sub operation = iota
	add
	mul
)

// 测试加减乘的溢出
func TestOverflow(t *testing.T) {
	for i, test := range []struct {
		x        uint64
		y        uint64
		overflow bool
		op       operation
	}{
		// add operations
		{math.MaxUint64, 1, true, add},
		{math.MaxUint64 - 1, 1, false, add},

		// sub operations
		{0, 1, true, sub},
		{0, 0, false, sub},

		// mul operations
		{0, 0, false, mul},
		{10, 10, false, mul},
		{math.MaxUint64, 2, true, mul},
		{math.MaxUint64, 1, false, mul},
	} {
		var overflows bool
		switch test.op {
		case sub:
			_, overflows = math.SafeSub(test.x, test.y)
		case add:
			_, overflows = math.SafeAdd(test.x, test.y)
		case mul:
			_, overflows = math.SafeMul(test.x, test.y)
		}

		if test.overflow != overflows {
			t.Errorf("%d failed. Expected test to be %v, got %v", i, test.overflow, overflows)
		}
		fmt.Println(test.op, test.x, test.y, overflows)
		fmt.Println("----")
	}
}

// 16进制转换成math.HexOrDecimal64，其中的规则比较复杂
func TestHexOrDecimal64(t *testing.T) {
	tests := []struct {
		input string
		num   uint64
		ok    bool
	}{
		{"", 0, true},
		{"0", 0, true},
		{"0x0", 0, true},
		{"12345678", 12345678, true},
		{"0x12345678", 0x12345678, true},
		{"0X12345678", 0x12345678, true},
		// Tests for leading zero behaviour:
		{"0123456789", 123456789, true}, // note: not octal
		{"0x00", 0, true},
		{"0x012345678abc", 0x12345678abc, true},
		// Invalid syntax:
		{"abcdef", 0, false},
		{"0xgg", 0, false},
		// Doesn't fit into 64 bits:
		{"18446744073709551617", 0, false},
	}
	for _, test := range tests {
		var num math.HexOrDecimal64
		err := num.UnmarshalText([]byte(test.input))
		// if (err == nil) != test.ok {
		// 	t.Errorf("ParseUint64(%q) -> (err == nil) = %t, want %t", test.input, err == nil, test.ok)
		// 	continue
		// }
		// if err == nil && uint64(num) != test.num {
		// 	t.Errorf("ParseUint64(%q) -> %d, want %d", test.input, num, test.num)
		// }
		if err != nil {
			fmt.Println(test.input, num, "错误")
			fmt.Println("----")
			continue
		}
		fmt.Println(test.input, num)
		fmt.Println("----")
	}
}

// 字符串强制转uint64
func TestMustParseUint64(t *testing.T) {
	if v := math.MustParseUint64("12345"); v != 12345 {
		t.Errorf(`MustParseUint64("12345") = %d, want 12345`, v)
	}
	fmt.Println(math.MustParseUint64("12345"))
}

// 字符串强制转uint64，panic
func TestMustParseUint64Panic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("MustParseBig should've panicked")
		}
	}()
	math.MustParseUint64("ggg")
}
