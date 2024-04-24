// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Adapted from: https://go.dev/src/crypto/subtle/xor_test.go

package bitutil

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/bitutil"
)

// 按位异或
func TestXOR(t *testing.T) {

	p := make([]byte, 3)
	p[0] = 1
	p[1] = 2
	p[2] = 3
	q := make([]byte, 2)
	q[0] = 3
	q[1] = 4

	dst := make([]byte, 4)
	dst[2] = 22
	dst[3] = 33

	bitutil.XORBytes(dst, p, q) //dst的长度一定>=min(p的长度，q的长度)，否则会panic。dst的结果的长度是min(p的长度，q的长度)，剩余的是不变

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(dst)

}

// 按位与
func TestAND(t *testing.T) {
	p := make([]byte, 3)
	p[0] = 1
	p[1] = 2
	p[2] = 3
	q := make([]byte, 2)
	q[0] = 3
	q[1] = 4

	dst := make([]byte, 4)
	dst[2] = 22
	dst[3] = 33

	bitutil.ANDBytes(dst, p, q) //dst的长度一定>=min(p的长度，q的长度)，否则会panic。dst的结果的长度是min(p的长度，q的长度)，剩余的是不变

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(dst)
}

// 按位或
func TestOR(t *testing.T) {
	p := make([]byte, 3)
	p[0] = 1
	p[1] = 2
	p[2] = 3
	q := make([]byte, 2)
	q[0] = 3
	q[1] = 4

	dst := make([]byte, 4)
	dst[2] = 22
	dst[3] = 33

	bitutil.ORBytes(dst, p, q) //dst的长度一定>=min(p的长度，q的长度)，否则会panic。dst的结果的长度是min(p的长度，q的长度)，剩余的是不变

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(dst)
}

// 测试一个字节切片p中是否至少有一个位被设置（即值为1）
func TestTest(t *testing.T) {
	p := make([]byte, 3)
	r := bitutil.TestBytes(p) //p全0返回false
	fmt.Println(r)

	p[0] = 2
	p[1] = 2
	p[2] = 2
	r = bitutil.TestBytes(p) //p不全0返回true
	fmt.Println(r)

}
