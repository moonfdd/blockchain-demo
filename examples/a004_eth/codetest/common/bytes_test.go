// Copyright 2014 The go-ethereum Authors
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

package common

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// 复制[]byte
func TestCopyBytes(t *testing.T) {
	if false {
		common.Report("1", 2, "3.0") //还会打印堆栈信息
		return
	}
	if false {
		common.PrintDeprecationWarning("哈哈") // 这是一个Go语言的函数定义，名为PrintDeprecationWarning。这个函数的目的是打印一个包含给定字符串的警告框，通常用于提醒用户某个功能或特性已被弃用。
		return
	}
	input := []byte{1, 2, 3, 4}

	v := common.CopyBytes(input)
	if !bytes.Equal(v, []byte{1, 2, 3, 4}) {
		t.Fatal("not equal after copy")
	}
	fmt.Println(input, v)
	v[0] = 99
	if bytes.Equal(v, input) {
		t.Fatal("result is not a copy")
	}
	fmt.Println(input, v)
}

// 左边补0
func TestLeftPadBytes(t *testing.T) {
	val := []byte{0, 1, 2, 3, 4}
	padded := []byte{0, 0, 0, 0, 1, 2, 3, 4}
	var r []byte

	r = common.LeftPadBytes(val, 7) //不足7位，补0
	fmt.Println(val, padded, r)
	r = common.LeftPadBytes(val, 2)
	fmt.Println(val, padded, r)
}

// 右边补0
func TestRightPadBytes(t *testing.T) {
	val := []byte{1, 2, 3, 4}
	padded := []byte{1, 2, 3, 4, 0, 0, 0, 0}

	if r := common.RightPadBytes(val, 8); !bytes.Equal(r, padded) {
		t.Fatalf("RightPadBytes(%v, 8) == %v", val, r)
	}
	if r := common.RightPadBytes(val, 2); !bytes.Equal(r, val) {
		t.Fatalf("RightPadBytes(%v, 2) == %v", val, r)
	}
}

// 16进制转[]byte，可以不含0x，可以是奇数位
func TestFromHex(t *testing.T) {
	input := "0xA"
	// expected := []byte{1}
	result := common.FromHex(input)
	// if !bytes.Equal(expected, result) {
	// 	t.Errorf("Expected %x got %x", expected, result)
	// }
	fmt.Println(result)
}

// 奇数位
func TestFromHexOddLength(t *testing.T) {
	input := "0x1"
	expected := []byte{1}
	result := common.FromHex(input)
	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %x got %x", expected, result)
	}
	fmt.Println(result)
}

// 不含前缀
func TestNoPrefixShortHexOddLength(t *testing.T) {
	input := "1"
	expected := []byte{1}
	result := common.FromHex(input)
	if !bytes.Equal(expected, result) {
		t.Errorf("Expected %x got %x", expected, result)
	}
	fmt.Println(result)
}

// 去掉右边0
func TestTrimRightZeroes(t *testing.T) {
	tests := []struct {
		arr []byte
		exp []byte
	}{
		{common.FromHex("0x00ffff00ff0000"), common.FromHex("0x00ffff00ff")},
		{common.FromHex("0x00000000000000"), []byte{}},
		{common.FromHex("0xff"), common.FromHex("0xff")},
		{[]byte{}, []byte{}},
		{common.FromHex("0x00ffffffffffff"), common.FromHex("0x00ffffffffffff")},
	}
	for i, test := range tests {
		got := common.TrimRightZeroes(test.arr) //// 去掉右边0
		if !bytes.Equal(got, test.exp) {
			t.Errorf("test %d, got %x exp %x", i, got, test.exp)
		}
		fmt.Println(test.arr, got)
		fmt.Println("----")
	}
}
