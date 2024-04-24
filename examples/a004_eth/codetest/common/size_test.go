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
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// 存储大小显示
func TestStorageSizeString(t *testing.T) {
	tests := []struct {
		size common.StorageSize
		str  string
	}{
		{2839274474874, "2.58 TiB"},
		{2458492810, "2.29 GiB"},
		{2381273, "2.27 MiB"},
		{2192, "2.14 KiB"},
		{12, "12.00 B"},
	}

	for _, test := range tests {
		if test.size.String() != test.str {
			t.Errorf("%f: got %q, want %q", float64(test.size), test.size.String(), test.str)
		}
		fmt.Println(test.size.TerminalString())
		fmt.Println("----")
	}
}

// 存储大小显示，日志或终端的函数
func TestStorageSizeTerminalString(t *testing.T) {
	tests := []struct {
		size common.StorageSize
		str  string
	}{
		{2839274474874, "2.58TiB"},
		{2458492810, "2.29GiB"},
		{2381273, "2.27MiB"},
		{2192, "2.14KiB"},
		{12, "12.00B"},
	}

	for _, test := range tests {
		if test.size.TerminalString() != test.str {
			t.Errorf("%f: got %q, want %q", float64(test.size), test.size.TerminalString(), test.str)
		}
		fmt.Println(test.size.TerminalString())
		fmt.Println("----")
	}
}
